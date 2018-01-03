/*
 *    Copyright 2017 Hans Mündelein
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package kubevaulter

import (
	vault "github.com/hashicorp/vault/api"
)

type ApiWrapper struct {
	loginForge LoginForge
	client     *vault.Client
	api        *vault.Logical
}

func NewApiWrapper(loginForger LoginForge, addr string) (*ApiWrapper, error) {
	config := vault.DefaultConfig()
	config.Address = addr
	vc, err := vault.NewClient(config)
	if err != nil {
		return nil,err
	}
	logical := vc.Logical()
	aw  := ApiWrapper{loginForge: loginForger,client:vc,api:logical}
	return &aw, nil
}

func (aw *ApiWrapper) KubeAuth() (*vault.Secret, error) {
	secret, err := aw.api.Write(aw.loginForge.GetPath(),aw.loginForge.ForgeRequest())
	if err != nil {
		return nil,err
	}
	aw.client.SetToken(secret.Auth.ClientToken)
	return secret, nil
}

func (aw *ApiWrapper) Read(path string) (*vault.Secret, error) {
	resp, err :=aw.api.Read(path)

	if err != nil {
		return nil,err
	}
	return  resp, nil
}

