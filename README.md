# ms-graph-auth-helper

一个可以快速帮助获取Microsoft Graph API令牌的小工具

并且可以自定义获取权限的类型

~~在弄网盘API的时候觉得太麻烦了就搞了个这个~~

![run](https://user-images.githubusercontent.com/7535224/147194336-5c75b46f-bd1d-4b71-8587-a09a0cdf1005.png)

![res](https://user-images.githubusercontent.com/7535224/147194576-7222b3f5-813d-4da1-86ed-69b60f09eb3c.png)

## 如何使用

**以获取OneDriver相关的API权限为例**

在获取令牌时，请按照以下步骤操作

#### 注册新应用

- 打开[Azure应用管理面板](https://portal.azure.com/#blade/Microsoft_AAD_RegisteredApps/ApplicationsListBlade)注册新的应用

- **受支持的帐户类型** 选项选择 **任何组织目录(任何 Azure AD 目录 - 多租户)中的帐户和个人 Microsoft 帐户(例如，Skype、Xbox)**

- 重定向URL（可选）中选择 **Web** 类型并将URL值设置为 **`http://localhost:11451/auth`**

#### 权限

- 切换至 **API权限** 选项卡点击 **添加权限**

- 选择 **Microsoft Graph** 分类

- 使用搜索功能找到并勾选以下权限：
  
  - Files.Read
  
  - Files.Read.All
  
  - Files.ReadWrite
  
  - Files.ReadWrite.All
  
  - offline_access

- 点击 **添加权限** 完成授权

#### 应用程序(客户端) ID [Client Id]

- 切换至 **概述** 选项卡复制**应用程序(客户端) ID** 并**自行保管**

#### 机密值[Client Secret]

- 切换至 **证书和密码** 选项卡

- 点击 **新建客户端密码** 并创建一个新的密钥

- 创建完成后复制密钥的值并**自行保管**

#### 启动工具按照提示完成令牌的获取
