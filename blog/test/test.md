好的，你可以按照以下步骤搭建一个类似 `username.github.io` 的博客：

1. 注册一个 Github 账号，假设你的账号名为 `yourusername`。
2. 创建一个名为 `yourusername.github.io` 的仓库。
3. 安装 Hugo，可以使用以下命令：

```
Copy code
brew install hugo
```

1. 创建 Hugo 站点，使用以下命令：

```
arduinoCopy code
hugo new site yourusername.github.io
```

1. 在 Hugo 站点目录下，使用以下命令创建一个新主题：

```
bashCopy code
cd yourusername.github.io
git init
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
echo 'theme = "ananke"' >> config.toml
```

1. 在 Hugo 站点目录下创建一个新的 Git 仓库，用于保存博客源代码：

```
sqlCopy code
git init
git remote add origin git@github.com:yourusername/yourusername.github.io.git
git add .
git commit -m "Initial commit"
git push -u origin master
```

1. 创建一个新的 Git 分支用于保存博客生成的 HTML 文件：

```
csharpCopy code
git checkout -b gh-pages
git add -f public/
git commit -m "Initial gh-pages commit"
git push -u origin gh-pages
```

1. 创建一个名为 `.github/workflows/deploy.yml` 的文件，用于配置 GitHub Actions 自动部署：

```yaml
yamlCopy code
name: Deploy

on:
  push:
    branches:
      - master

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Setup Hugo
      uses: peaceiris/actions-hugo@v2
      with:
        hugo-version: "0.83.1"
    - name: Build
      run: hugo --minify
    - name: Deploy
      uses: peaceiris/actions-gh-pages@v3
      with:
        personal_token: ${{ secrets.GH_PAGES_TOKEN }}
        publish_dir: ./public
```

1. 生成一个 GitHub Personal Access Token，用于 GitHub Actions 自动部署，可以按照以下步骤生成：

- 登录 GitHub 账号，点击