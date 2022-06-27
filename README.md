# goApplet

以go-frame框架为基，搭建的一个基础的聊天工具

This repo implements some demos using GoFrame.
1. A simple websocket chat service.
1. Basic API example for user SignUp/SignIn.
1. Universal CURD service.

## Installation

### 1. You need a go development environment setup before everything starts taking off.
### 2. Use `git clone` cloing the repo to your local folder.
```
git clone 
```

### 3. Import `document/sql/create.sql` to your database.

### 4. Create configuration file from `config.example.toml`.
```
cp config/config.example.toml config/config.toml
```
Update `config.toml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:
```
  SERVER  | DOMAIN  | ADDRESS | METHOD |        ROUTE        |                      HANDLER                       |                                                  MIDDLEWARE                                                    
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat               | goApplet/app/api.(*chatApi).Index                  | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/index         | goApplet/app/api.(*chatApi).Index                  | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/setname       | goApplet/app/api.(*chatApi).SetName                | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/websocket     | goApplet/app/api.(*chatApi).WebSocket              | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /swagger/*          | github.com/gogf/swagger.(*Swagger).Install.func1.1 | HOOK_BEFORE_SERVE                                                                                              
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/checknickname | goApplet/app/api.(*userApi).CheckNickName          | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/checkpassport | goApplet/app/api.(*userApi).CheckPassport          | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/issignedin    | goApplet/app/api.(*userApi).IsSignedIn             | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/profile       | goApplet/app/api.(*userApi).Profile-fm             | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm,service.(*middlewareService).Auth-fm  
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signin        | goApplet/app/api.(*userApi).SignIn                 | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signout       | goApplet/app/api.(*userApi).SignOut                | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signup        | goApplet/app/api.(*userApi).SignUp                 | service.(*middlewareService).Ctx-fm,service.(*middlewareService).CORS-fm                                       
----------|---------|---------|--------|---------------------|----------------------------------------------------|----------------------------------------------------------------------------------------------------------------
```

# GoFrame Sites
### GoFrame Repository
* [https://github.com/gogf/gf](https://github.com/gogf/gf)
* [https://gitee.com/johng/gf](https://gitee.com/johng/gf)

### GoFrame Home
* [https://goframe.org](https://goframe.org)
