# 朱雀-nodejs发布系统
朱雀发布系统是专门为nodejs发布而开发的系统，使用go语言开发，前端使用了layui mini框架，简单易上手。
朱雀发布系统前后端分离，但又是一体的，无需部署两套服务。
理论上朱雀发布系统可以发布其他语言应用程序，但其与nodejs更加相得益彰。
原因是朱雀发布系统依赖于PM2，PM2大家都知道，与nodejs几乎是绝配，所以有nodejs的地方很大可能有PM2，所以使用朱雀发布系统就显得更加简单了，无需专门安装PM2。

数据库使用sqlite，无需单独安装和配置。

# 使用框架文档
[前端框架](http://layuimini.99php.cn/docs/index.html)
[前端框架](http://layuimini.99php.cn/onepage/v2/index.html)
[前端框架](https://www.layui.com/doc/)
[图标库](https://fontawesome.dashgame.com/)
[后端框架gorm](http://gorm.book.jasperxu.com/)
[后端框架gin](https://github.com/gin-gonic/gin#using-middleware)

## 部署流程
1. 在服务器指定位置下载源码。
2. 安装go环境。go可以交叉编译，但是由于sqlite的缘故，windows环境下并不能顺利的编译linux版本，所以最好还是在linux环境下编译。
3. 配置`conf.yaml`文件，参照`conf-sample.yaml`文件。
4. 项目目录中编译`go build`，第一次会安装依赖会慢一些。
5. `./zhuque`启动服务。

## 权限架构
该系统权限使用了基于角色的访问控制方法（RBAC）。是目前公认的解决大型企业的统一资源访问控制的有效方法。 其显著的两大特征是：1.减小授权管理的复杂性，降低管理开销。2.灵活地支持企业的安全策略，并对企业的变化有很大的伸缩性。

参考文档：

[官方原文](https://pm2.keymetrics.io/docs/usage/deployment/)

[PM2自动部署代码流程总结](https://segmentfault.com/a/1190000017310047)

[pm2 官方文档 学习笔记](https://my.oschina.net/u/4400196/blog/3283439)

[通过Github与PM2部署Node应用](https://zhuanlan.zhihu.com/p/20940096)

[deploy](https://github.com/Unitech/PM2/blob/0.14.7/ADVANCED_README.md#deployment-options)

## [Getting started](#getting-started)

PM2 embeds a simple and powerful deployment system with revision tracing.

Please read the [Considerations to use PM2 deploy](#considerations).

## [Simple deploy](#simple-deploy)

You only need to add a “deploy” attribute to the ecosystem.json file. Here is a bare minimum to deploy an application:

process.json:

```
{
   "apps" : [{
      "name" : "HTTP-API",
      "script" : "http.js"
   }],
   "deploy" : {
     // "production" is the environment name
     "production" : {
       "user" : "ubuntu",
       "host" : ["192.168.0.13"],
       "ref"  : "origin/master",
       "repo" : "git@github.com:Username/repository.git",
       "path" : "/var/www/my-repository",
       "post-deploy" : "npm install; grunt dist"
      },
   }
}
```

/bin/bash:
```
# Setup deployment at remote location
$ pm2 deploy production setup

# Update remote version
$ pm2 deploy production update

# Revert to -1 deployment
$ pm2 deploy production revert 1

# execute command on remote machines
$ pm2 deploy production exec "pm2 reload all"
```

## [Complete tutorial](#complete-tutorial)

1- Generate a sample ecosystem.json file that lists the processes and the deployment environment.

`pm2 ecosystem`

In the current folder a `ecosystem.json` file will be created.
    It contains this:
```
{
  // Applications part
  "apps" : [{
    "name"      : "API",
    "script"    : "app.js",
    "env": {
      "COMMON_VARIABLE": "true"
    },
    // Environment variables injected when starting with --env production
    // http://pm2.keymetrics.io/docs/usage/application-declaration/#switching-to-different-environments
    "env_production" : {
      "NODE_ENV": "production"
    }
  },{
    "name"      : "WEB",
    "script"    : "web.js"
  }],
  // 部署部分
  // Here you describe each environment
  "deploy" : {
    "production" : {
      "user" : "node",
      // 服务器集群
      "host" : ["212.83.163.1", "212.83.163.2", "212.83.163.3"],
      // 分支
      "ref"  : "origin/master",
      // Git 地址
      "repo" : "git@github.com:repo.git",
      // 应用在服务器上的地址
      "path" : "/var/www/production",
      // Can be used to give options in the format used in the configura-
      // tion file.  This is useful for specifying options for which there
      // is no separate command-line flag, see 'man ssh'
      // can be either a single string or an array of strings
      "ssh_options": "StrictHostKeyChecking=no",
      // To prepare the host by installing required software (eg: git)
      // even before the setup process starts
      // can be multiple commands separated by the character ";"
      // or path to a script on your local machine
      "pre-setup" : "apt-get install git",
      // Commands / path to a script on the host machine
      // This will be executed on the host after cloning the repository
      // eg: placing configurations in the shared dir etc
      "post-setup": "ls -la",
      // Commands to execute locally (on the same machine you deploy things)
      // Can be multiple commands separated by the character ";"
      "pre-deploy-local" : "echo 'This is a local executed command'"
      // Commands to be executed on the server after the repo has been cloned
      "post-deploy" : "npm install && pm2 startOrRestart ecosystem.json --env production"
      // Environment variables that must be injected in all applications on this env
      "env"  : {
        "NODE_ENV": "production"
      }
    },
    "staging" : {
      "user" : "node",
      "host" : "212.83.163.1",
      "ref"  : "origin/master",
      "repo" : "git@github.com:repo.git",
      "path" : "/var/www/development",
      "ssh_options": ["StrictHostKeyChecking=no", "PasswordAuthentication=no"],
      "post-deploy" : "pm2 startOrRestart ecosystem.json --env dev",
      "env"  : {
        "NODE_ENV": "staging"
      }
    }
  }
}
```

Edit the file according to your needs.

2- Be sure that you have the public ssh key on your local machine
```
ssh-keygen -t rsa
ssh-copy-id node@myserver.com
```

If you encounter any errors, please refer to the troubleshooting section below.

3- Now initialize the remote folder with:

`pm2 deploy <configuration_file> <environment> setup`

Example:

`pm2 deploy ecosystem.json production setup`

This command will create the folders on your remote server.

4- Deploy your code

`pm2 deploy ecosystem.json production`

Now your code will be populated, installed and started with PM2.

## [Deployment options](#deployment-options)

Display deploy help via `pm2 deploy help`:

```
pm2 deploy <configuration_file> <environment> <command>

  Commands:
    setup                run remote setup commands
    update               update deploy to the latest release
    revert [n]           revert to [n]th last deployment or 1
    curr[ent]            output current release commit
    prev[ious]           output previous release commit
    exec|run <cmd>       execute the given <cmd>
    list                 list previous deploy commits
    [ref]                deploy to [ref], the "ref" setting, or latest tag
```

## [Use different set of env variables](#use-different-set-of-env-variables)

In the `post-deploy` attribute, you may have noticed the command `pm2 startOrRestart ecosystem.json --env production`. The `--env <environment_name>` allows to inject different sets of environment variables.

Read more [here](http://pm2.keymetrics.io/docs/usage/application-declaration/#switching-to-different-environments).

## [Related Commands](#related-commands)

```
pm2 startOrRestart all.json            # Invoke restart on all apps in JSON
pm2 startOrReload all.json             # Invoke reload
```

## [Multi host deployment](#multi-host-deployment)

To deploy to multiple hosts in the same time, you just have to declare each host in an array under the attribute `host`.
```
{
  [...]
  "deploy" : {
    "production" : {
      "user" : "node",
      // Multi host in a js array
      "host" : ["212.83.163.1", "212.83.163.2", "212.83.163.3"],
      "ref"  : "origin/master",
      "repo" : "git@github.com:repo.git",
      "path" : "/var/www/production",
      "pre-setup" : "echo 'commands or local script path to be run on the host before the setup process starts'",
      "post-setup": "echo 'commands or a script path to be run on the host after cloning the repo'",
      "post-deploy" : "pm2 startOrRestart ecosystem.json --env production",
      "pre-deploy-local" : "echo 'This is a local executed command'"
    }
  [...]
}
```

## [Using SSH keys](#using-ssh-keys)

You just have to add the “key” attribute with path to the public key, see below example :
```
    "production" : {
      "key"  : "/path/to/some.pem", // path to the public key to authenticate
      "user" : "node",              // user used to authenticate
      "host" : "212.83.163.1",      // where to connect
      "ref"  : "origin/master",
      "repo" : "git@github.com:repo.git",
      "path" : "/var/www/production",
      "post-deploy" : "pm2 startOrRestart ecosystem.json --env production"
    },
```
## [Force deployment](#force-deployment)

You may get this message:

```
--> Deploying to dev environment
--> on host 192.168.1.XX

  push your changes before deploying

Deploy failed
```
That means that you have changes in your local system that aren’t pushed inside your git repository, and since the deploy script get the update via `git pull` they will not be on your server. If you want to deploy without pushing any data, you can append the `--force` option:

`pm2 deploy ecosystem.json production --force`

## [Considerations](#considerations)

*   You can use the option `--force` to skip local change detection
*   You might want to commit your node_modules folder ([#622](https://github.com/Unitech/pm2/issues/622)) or add the `npm install` command to the `post-deploy` section: `"post-deploy" : "npm install && pm2 startOrRestart ecosystem.json --env production"`
*   Verify that your remote server has the permission to git clone the repository
*   You can declare specific environment variables depending on the environment you want to deploy the code to. For instance to declare variables for the production environment, add “env_production”: {} and declare the variables.
*   By default, PM2 will use `ecosystem.json`. So you can skip the <configuration_file> options if this is the case</configuration_file>
*   You can embed the “apps” &amp; “deploy” section in the package.json
*   It deploys your code via ssh, you don’t need any dependencies
*   Processes are initialized / started automatically depending on the application name in `ecosystem.json`
*   PM2-deploy repository can be found here: [pm2-deploy](https://github.com/Unitech/pm2-deploy)
*   **WINDOWS** : see point below (at the end)

## [Troubleshooting](#troubleshooting)

##### SSH clone errors

In most cases, these errors will be caused by `pm2` not having the correct keys to clone your repository. You need to verify at every step that the keys are available.

**Step 1**
If you are certain your keys are correctly working, first try running `git clone your_repo.git` on the target server. If it succeeds, move onto the next steps. If it failed, make sure your keys are stored both on the server and on your git account.

**Step 2**
By default `ssh-copy-id` copies the default identiy, usually named `id_rsa`. If that is not the appropriate key:

`ssh-copy-id -i path/to/my/key your_username@server.com`

This adds your public key to the `~/.ssh/authorized_keys` file.

**Step 3**
If you get the following error:
```
--> Deploying to production environment
--> on host mysite.com
  ○ hook pre-setup
  ○ running setup
  ○ cloning git@github.com:user/repo.git
Cloning into '/var/www/app/source'...
Permission denied (publickey).
fatal: Could not read from remote repository.

Please make sure you have the correct access rights and that the repository exists.

**Failed to clone**

Deploy failed
```

…you may want to create a ssh config file. This is a sure way to ensure that the correct ssh keys are used for any given repository you’re trying to clone. See [this example](https://gist.github.com/Protosac/c3fb459b1a942f161f23556f61a67d66):

```
# ~/.ssh/config
Host alias
    HostName myserver.com
    User username
    IdentityFile ~/.ssh/mykey
# Usage: `ssh alias`
# Alternative: `ssh -i ~/.ssh/mykey username@myserver.com`

Host deployment
    HostName github.com
    User username
    IdentityFile ~/.ssh/github_rsa
# Usage:
# git@deployment:username/anyrepo.git
# This is for cloning any repo that uses that IdentityFile. This is a good way to make sure that your remote cloning commands use the appropriate key
```

## [Windows Consideration](#windows-consideration)

To run the deploy script under Windows, you need to use a unix shell like bash, so we recommend to install either [Git bash](https://git-scm.com/download/win), [Babun](http://babun.github.io/) or  [Cygwin](https://cygwin.com/install.html)

## [Contributing](#contributing)

The module is [https://github.com/Unitech/pm2-deploy](https://github.com/Unitech/pm2-deploy)
Feel free to PR for any changes or fix.
