# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [0.7.2](https://github.com/sinlov/filebrowser-client/compare/0.7.1...v0.7.2) (2024-04-26)

### 🐛 Bug Fixes

* change `ResourceDownload` support `override` true to mkdir as local ([58a64664](https://github.com/sinlov/filebrowser-client/commit/58a646644aacb4fca7c5c6e15e324908cbade7d2)), fix [#5](https://github.com/sinlov/filebrowser-client/issues/5)

## [0.7.1](https://github.com/sinlov/filebrowser-client/compare/0.7.0...v0.7.1) (2024-04-26)

### 🐛 Bug Fixes

* `sendSaveFile` print curl without debug ([5cfa026b](https://github.com/sinlov/filebrowser-client/commit/5cfa026b7fbb8361864c589f7e264590318db710)), fix [#3](https://github.com/sinlov/filebrowser-client/issues/3)

## [0.7.0](https://github.com/sinlov/filebrowser-client/compare/0.6.1...v0.7.0) (2024-04-24)

### ✨ Features

* add `file_browser_log` package for management log ([d6872a28](https://github.com/sinlov/filebrowser-client/commit/d6872a282ebabbd9f7982f03733c539d55c70702))

* improve the code structure for easy maintenance ([70982691](https://github.com/sinlov/filebrowser-client/commit/709826918d5e8882c87765b93ce447587f871b5e)), feat [#1](https://github.com/sinlov/filebrowser-client/issues/1)

### ♻ Refactor

* use `github.com/sinlov-go/unittest-kit` for unit test of this project ([0d3e112a](https://github.com/sinlov/filebrowser-client/commit/0d3e112af4255995c3c5a047885c76b3d805316b))

### 👷‍ Build System

* github.com/urfave/cli/v2 v2.4.1 support go 1.11 then v2.4.2 update to go 1.18 ([5670287d](https://github.com/sinlov/filebrowser-client/commit/5670287dc7eb56c009b3f290ed9936f356f7868e))

## [0.6.1](https://github.com/sinlov/filebrowser-client/compare/0.6.0...v0.6.1) (2024-03-06)

### 📝 Documentation

* update test caes and doc of useage ([64465fff](https://github.com/sinlov/filebrowser-client/commit/64465fff857f1229ccc6014e92fd1a493a740570))

## [0.6.0](https://github.com/sinlov/filebrowser-client/compare/0.5.0...v0.6.0) (2024-03-06)

### ✨ Features

* update full build kit for golang ([71da0727](https://github.com/sinlov/filebrowser-client/commit/71da0727946bd6daf1a508a62de6b28007fd97df))

### 👷‍ Build System

* add .gitattributes ([f2002535](https://github.com/sinlov/filebrowser-client/commit/f20025357e79d829e06b7ca18818f1e20c4ac89a))

## [0.5.0](https://github.com/sinlov/filebrowser-client/compare/v0.4.0...v0.5.0) (2023-02-04)

### Features

* let SharePost support Infinite when expires set 0 ([803a125](https://github.com/sinlov/filebrowser-client/commit/803a12515f0368643c0f43232932fd64c02d72cb))

## [0.5.0](https://github.com/sinlov/filebrowser-client/compare/v0.4.0...v0.5.0) (2023-02-04)

### Features

* let SharePost support Infinite when expires set 0 ([803a125](https://github.com/sinlov/filebrowser-client/commit/803a12515f0368643c0f43232932fd64c02d72cb))

## [0.4.0](https://github.com/sinlov/filebrowser-client/compare/v0.3.0...v0.4.0) (2023-02-04)

### Features

* add ResourcesPostFile file sha256 check ([abea935](https://github.com/sinlov/filebrowser-client/commit/abea935af5c22233027125488ed7c7f7bbf00267))

## [0.3.0](https://github.com/sinlov/filebrowser-client/compare/v0.2.1...v0.3.0) (2023-02-04)

### Features

* add url check by path ([0cdcd10](https://github.com/sinlov/filebrowser-client/commit/0cdcd10dec57060fff6bb6a8f208c2314f917cf1))

* mark version 0.3.0 ([1cbd924](https://github.com/sinlov/filebrowser-client/commit/1cbd9245d092a365a0ac976d7f6e9cf5ffd94af7))

* try to fix path error at different OS ([4930735](https://github.com/sinlov/filebrowser-client/commit/49307354e139d85436940051f4346de82d5441fb))

### Bug Fixes

* fix send file error at windows path ([6742397](https://github.com/sinlov/filebrowser-client/commit/67423971d28b4398d1816ed6a8918e8ae1036df5))

### [0.2.1](https://github.com/sinlov/filebrowser-client/compare/v0.1.3...v0.2.1) (2023-02-03)

### Features

* add folder.WalkAllByGlob function to walk path by glob ([b5dc5d5](https://github.com/sinlov/filebrowser-client/commit/b5dc5d5d66b9bc72db752b04e0290e10939d2414))

* add tools.StrArrRemoveDuplicates and test case Benchmark test ([486868d](https://github.com/sinlov/filebrowser-client/commit/486868da41f8aa994c87eb2b56477c2ff199fd7f))

## <small>0.1.3 (2023-02-01)</small>

* feat: change private func sendRespRaw has showCurl to close send file log ([65f25c5](https://github.com/sinlov/filebrowser-client/commit/65f25c5))

* test: add some unit test case ([60d2eda](https://github.com/sinlov/filebrowser-client/commit/60d2eda))

* docs: add help of mod require ([46332bd](https://github.com/sinlov/filebrowser-client/commit/46332bd))

* docs: update help of version use ([4766c6b](https://github.com/sinlov/filebrowser-client/commit/4766c6b))

## <small>0.1.2 (2023-01-31)</small>

* feat: add github.com/urfave/cli/v2 for cli ([80327a4](https://github.com/sinlov/filebrowser-client/commit/80327a4))

* feat: add main for base cli, then update Api FileBrowserClient ([a193c49](https://github.com/sinlov/filebrowser-client/commit/a193c49))

* feat: add ResourcesPostDirectoryFiles ResourceDownload, change SharePost will show more info ([fa0f635](https://github.com/sinlov/filebrowser-client/commit/fa0f635))

* feat: add ResourcesPostOne api and test case ([387987a](https://github.com/sinlov/filebrowser-client/commit/387987a))

* feat: add SharePost and SharesGet ([5463346](https://github.com/sinlov/filebrowser-client/commit/5463346))

* feat: add tools/folder for post files ([0cc5ae0](https://github.com/sinlov/filebrowser-client/commit/0cc5ae0))

* feat: update coverage.txt for check ([e7e51a7](https://github.com/sinlov/filebrowser-client/commit/e7e51a7))

* ci: change go coverage test result update from local by local full env ([14d10bc](https://github.com/sinlov/filebrowser-client/commit/14d10bc))

* first commit ([073157d](https://github.com/sinlov/filebrowser-client/commit/073157d))
