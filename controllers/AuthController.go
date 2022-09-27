package controllers

import (
	dto "blog-api-golang/DTO"
	"blog-api-golang/models"
	"blog-api-golang/services"
	"blog-api-golang/utils"
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	userData := &dto.Register{}
	utils.ParseBody(r, userData)

	user := &services.User{}
	user = services.MatchUserProperties(*userData)

	user.Password, _ = utils.GenerateHashPassword(user.Password)
	u, db := services.CreateUser(user)
	var response dto.Response
	if u == nil {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username has been taken")
		response.Entity = nil
	} else {
		if db.Error == nil {
			response.Status = http.StatusOK
			response.Message = append(response.Message, "You are already registered!")
			response.Entity = u
		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to register")
			response.Message = append(response.Message, db.Error.Error())
			response.Entity = nil
		}

	}

	utils.EncodeJson(w, response, response.Status)

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	loginUser := &dto.Authentication{}
	utils.ParseBody(r, loginUser)
	authUser, exists := services.FindUserByUsername(loginUser.Username)
	if !exists {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username or Password is Incorrect")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	passwordIsMatched := utils.CheckPasswordHash(loginUser.Password, authUser.Password)

	if !passwordIsMatched {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username or Password is Incorrect")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	token, err := utils.GenerateJWT(authUser.Username)
	if err != nil {
		var response dto.Response
		response.Status = http.StatusInternalServerError
		response.Message = append(response.Message, "Something went wrong!")
		response.Message = append(response.Message, err.Error())
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	var tokenData dto.Token
	tokenData.Username = authUser.Username
	tokenData.Token = token
	var response dto.Response
	response.Status = http.StatusOK
	response.Message = append(response.Message, "You are already logged in")
	response.Entity = tokenData
	utils.EncodeJson(w, response, response.Status)

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var newProfile = &models.User{}
	utils.ParseBody(r, newProfile)
	var response dto.Response
	user, exits := services.FindUserByUsername(r.Header.Get("username"))
	tempSameUsername := true
	var sameUsername *bool = &tempSameUsername

	if exits {

		if &newProfile.Name != nil && newProfile.Name != "" {
			user.Name = newProfile.Name
			fmt.Println(user.Name)
		}

		if (newProfile.Username != "") && newProfile.Username != user.Username {
			tempSameUsername = false
			user.Username = newProfile.Username
			sameUsername = &tempSameUsername
		}

		_, dbres := services.UpdateUser(user)
		if dbres.Error == nil {
			if !*sameUsername {
				token, err := utils.GenerateJWT(user.Username)
				var tokenData dto.Token

				if err == nil {
					tokenData.Username = user.Username
					tokenData.Token = token
					response.Status = http.StatusOK
					response.Message = append(response.Message, "Profile successfully updated")
					response.Entity = tokenData
				} else {
					var response dto.Response
					response.Status = http.StatusInternalServerError
					response.Message = append(response.Message, "Something went wrong!")
					response.Message = append(response.Message, err.Error())
					response.Entity = nil
				}

			} else {
				response.Status = http.StatusOK
				response.Message = append(response.Message, "Profile successfully updated")
				response.Entity = nil
			}

		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to update profile")
			response.Message = append(response.Message, dbres.Error.Error())
			response.Entity = nil
		}

	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "User doesn't exist")
		response.Entity = nil
	}
	utils.EncodeJson(w, response, response.Status)

}
