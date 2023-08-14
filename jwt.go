package main

// Function to verify the OAuth2 token using JWT
// func verifyToken(token *oauth2.Token) (*jwt.Token, error) {
// 	// Use GitHub's OAuth2 token verification endpoint
// 	jwksURL := "https://github.com/login/oauth/verify"

// 	// Fetch the public keys from GitHub to verify the token
// 	jwks, err := jwt.FetchHTTP(jwksURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch public keys: %v", err)
// 	}

// 	// Use the public keys to verify the token
// 	verifiedToken, err := jwt.ParseWithClaims(token.AccessToken, &jwt.StandardClaims{}, jwks.Key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	return verifiedToken, nil
// }
