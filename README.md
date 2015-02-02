# i2med-server
## 用户接口

### POST /user/login 


|param_name|type|comment|
:-------------: | :---: | :-------: |
|user|string|用户名|
|pwd|string|密码|


### POST /user/register
|param_name|type|comment|
:------------: | :---: | :--------: |
|name|string|用户名|
|pwd|string|密码，加密过|
|age|int|年龄|
|title|string|职称|
|department|string|部门|
|province|string|省|
|city|string|城市|


## 朋友接口

### POST /friends/add/    
|param_name|type|comment|
:-------------: | :---: | :-------: |
|userId|int|用户ID|
|friendId|int|朋友用户ID|          

### POST /friends/list/
|param_name|type|comment|
:-------------: | :---: | :-------: |
|userId|int|用户ID|
             

### POST /friends/get/              
|param_name|type|comment|
:-------------: | :---: | :-------: |
|userId|int|用户ID|
|friendId|int|朋友用户ID|


## 状态接口

### 发状态 POST /message/post
|param_name|type|comment|
:-------------: | :---: | :-------: |
|userId|int|用户ID|
|message|string|消息文本|