# go-github-reactions

## 준비하기

github 접근용 token을 발급한다.
1. `https://{GITHUB_HOST}/settings/tokens/new` > Generate New Token 
2. Check `repo`, `users`
3. Click `Generate token` button

## 설정하기

`.env`파일을 생성한 후 아래와 같이 입력한다.

```
GITHUB_API=https://api.github.com
GITHUB_TOKEN={your token}
```

## 실행하기

```
$ npm install && npm run start -- {owner} {repo} {issueNumber}
```

ex)
```
// 모든 리액션 가져오기
$ npm install && npm run start -- Microsoft vscode 164
```

```
// +1 리액션 가져오기
$ npm install && npm run start -- Microsoft vscode 164 +1
```

## 리액션 종류

|  content  |	  emoji  |
|-----------|----------|
|       `+1`|      :+1:|
|       `-1`|      :-1:|
|    `laugh`|   :smile:|
| `confused`|:confused:|
|    `heart`|   :heart:|
|   `hooray`|    :tada:|
