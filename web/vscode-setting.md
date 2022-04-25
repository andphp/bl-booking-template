### npm install eslint -g

```json
{
    "workbench.colorTheme": "Default Dark+",
    "tabnine.experimentalAutoImports": true,
    "workbench.iconTheme": "vscode-icons",
    "go.autocompleteUnimportedPackages": true,
    "go.enableCodeLens": {
        "runtest": false
    },
    "go.formatTool": "goimports",
    "window.zoomLevel": 1,

    // react
    "settings": {
        "liveServer.settings.AdvanceCustomBrowserCmdLine": "",
        "liveServer.settings.CustomBrowser": "chrome",
        "editor.fontsize": 16,
        "editor.formatOnSave": true
      },
      "editor.tabSize": 2,
      "[javascriptreact]": {
        "editor.defaultFormatter": "vscode.typescript-language-features"
      },
      "[javascript]": {
        "editor.defaultFormatter": "vscode.typescript-language-features"
      },
      "editor.fontSize": 16,
      "[html]": {
        "editor.defaultFormatter": "vscode.html-language-features"
      },
      "[json]": {
        "editor.defaultFormatter": "esbenp.prettier-vscode"
      },
      "editor.suggest.snippetsPreventQuickSuggestions": false,
      "[vue]": {
        "editor.defaultFormatter": "octref.vetur"
      },
      "editor.minimap.enabled": false,
      "workbench.activityBar.visible": true,
      "standard.autoFixOnSave": true,
      "emmet.syntaxProfiles": {
        "vue-html": "html",
        "vue": "html"
      },
      "emmet.triggerExpansionOnTab": true,
      "[jsonc]": {
        "editor.defaultFormatter": "esbenp.prettier-vscode"
      },
      "editor.formatOnType": true,
      "editor.formatOnSave": true,
      //配置eslint
      "editor.codeActionsOnSave": {
        "source.fixAll.eslint": true
      },
      "eslint.format.enable": true,
      //autoFix默认开启，只需输入字符串数组即可
      "eslint.validate": ["javascript", "vue", "html"],
      "launch": {
        "configurations": [],
        "compounds": []
      },
      "vetur.format.defaultFormatterOptions": {
        "prettier": {
          "semi": false,
          "singleQuote": true,
          "printWidth": 200,
          "trailingComma": "none"
        }
      },
      "javascript.format.insertSpaceBeforeFunctionParenthesis": false,
      "emmet.includeLanguages": {
        "javascript": "javascriptreact"
      },
}
```