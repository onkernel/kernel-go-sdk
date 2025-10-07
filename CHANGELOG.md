# Changelog

## 0.14.1 (2025-10-07)

Full Changelog: [v0.14.0...v0.14.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.14.0...v0.14.1)

### Features

* WIP browser extensions ([22323cd](https://github.com/onkernel/kernel-go-sdk/commit/22323cdf2376bf2016cf1ebd1384f4c1b5ca6752))

## 0.14.0 (2025-10-03)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.13.0...v0.14.0)

### Features

* Http proxy ([0ebe5f1](https://github.com/onkernel/kernel-go-sdk/commit/0ebe5f1f007fd53267ac096e044f9fd513c28d7b))

## 0.13.0 (2025-10-01)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.12.0...v0.13.0)

### Features

* Update oAPI and data model for proxy status ([87b8365](https://github.com/onkernel/kernel-go-sdk/commit/87b8365a941544eaf38b09ca66d6f9a97e3b050b))

## 0.12.0 (2025-09-30)

Full Changelog: [v0.11.5...v0.12.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.5...v0.12.0)

### Features

* Return proxy ID in browsers response ([711c52f](https://github.com/onkernel/kernel-go-sdk/commit/711c52facea0f2170938a0a68040c6781f2c19ee))

## 0.11.5 (2025-09-29)

Full Changelog: [v0.11.4...v0.11.5](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.4...v0.11.5)

### Features

* Add App Version to Invocation and add filtering on App Version ([81f87e2](https://github.com/onkernel/kernel-go-sdk/commit/81f87e27ae242be9f7cfc75a6147e9eac669d4c4))
* Fix my incorrect grammer ([f04186f](https://github.com/onkernel/kernel-go-sdk/commit/f04186fd786d47e49c42028c0237e02fd08d03b1))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([32304ba](https://github.com/onkernel/kernel-go-sdk/commit/32304baca3b01de677fd705da2b03787b56fdf35))

## 0.11.4 (2025-09-25)

Full Changelog: [v0.11.3...v0.11.4](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.3...v0.11.4)

### Features

* getInvocations endpoint ([b9a983c](https://github.com/onkernel/kernel-go-sdk/commit/b9a983c1cf96d2aa6d22f57aaf98ded129e8d4d4))

## 0.11.3 (2025-09-24)

Full Changelog: [v0.11.2...v0.11.3](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.2...v0.11.3)

### Features

* Per Invocation Logs ([f1241f2](https://github.com/onkernel/kernel-go-sdk/commit/f1241f2d66a097bfb11a2cbd9fda9495b54b0690))

## 0.11.2 (2025-09-24)

Full Changelog: [v0.11.1...v0.11.2](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.1...v0.11.2)

### Features

* Add stainless CI ([ac377ec](https://github.com/onkernel/kernel-go-sdk/commit/ac377ecde22bf08aa41dbddf77cbfd686354c1ea))


### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([e893e50](https://github.com/onkernel/kernel-go-sdk/commit/e893e50418d4dcbd11b910f12fad79f53ca5eb85))


### Chores

* bump minimum go version to 1.22 ([8f9a2ac](https://github.com/onkernel/kernel-go-sdk/commit/8f9a2ac0e091826c5ed578e568a250b4f11cc05a))
* do not install brew dependencies in ./scripts/bootstrap by default ([856b62a](https://github.com/onkernel/kernel-go-sdk/commit/856b62a529e4bed1ae8205065dc6c0e6e50e5f6d))
* update more docs for 1.22 ([f38538f](https://github.com/onkernel/kernel-go-sdk/commit/f38538f4031a5d8045b949d6e9b06f5c76cecde9))

## 0.11.1 (2025-09-06)

Full Changelog: [v0.11.0...v0.11.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.11.0...v0.11.1)

### Features

* **api:** add pagination to the deployments endpoint ([9130f8c](https://github.com/onkernel/kernel-go-sdk/commit/9130f8c40783560bc02d5b494c60549dbe092dcd))
* **api:** pagination properties added to response (has_more, next_offset) ([65c5342](https://github.com/onkernel/kernel-go-sdk/commit/65c53429690169c3ca5068b94991af6eba3832ff))
* **api:** update API spec with pagination headers ([ad37eb2](https://github.com/onkernel/kernel-go-sdk/commit/ad37eb2bdc92a4145e5dba9d8ef8b4ad26c9b64a))


### Bug Fixes

* **client:** correctly convert header pagination value to int ([d42bd8a](https://github.com/onkernel/kernel-go-sdk/commit/d42bd8a79dc3f64ea6aa8c3f4239b8cfa8545d82))
* **internal:** unmarshal correctly when there are multiple discriminators ([649203d](https://github.com/onkernel/kernel-go-sdk/commit/649203d6621f6850ac4cb414b04b78c9693d1923))

## 0.11.0 (2025-09-04)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.10.0...v0.11.0)

### Features

* **api:** adding support for browser profiles ([481cdb3](https://github.com/onkernel/kernel-go-sdk/commit/481cdb3500744c9e4ec050e340a920302d8fea19))


### Bug Fixes

* close body before retrying ([a6a2e40](https://github.com/onkernel/kernel-go-sdk/commit/a6a2e4054c629d6ee85997ed81a1b14e70e594dc))


### Chores

* **internal:** codegen related update ([a7030ab](https://github.com/onkernel/kernel-go-sdk/commit/a7030abb99c06c675f60a4f2afde43d376d9981f))

## 0.10.0 (2025-08-27)

Full Changelog: [v0.9.1...v0.10.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.9.1...v0.10.0)

### Features

* **api:** new process, fs, and log endpoints ([fa85f19](https://github.com/onkernel/kernel-go-sdk/commit/fa85f19a8f4c664696d8839539a6ce6172a78e98))

## 0.9.1 (2025-08-15)

Full Changelog: [v0.9.0...v0.9.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.9.0...v0.9.1)

### Features

* **api:** add browser timeouts ([6c7b47f](https://github.com/onkernel/kernel-go-sdk/commit/6c7b47f69ccc4e12d9e21340c543d57b7fc6d314))

### Chores

* **internal:** codegen related update ([e4ca558](https://github.com/onkernel/kernel-go-sdk/commit/e4ca55843e4dbb9b7e71821ca58080a5bf25f025))
* **internal:** update comment in script ([9542333](https://github.com/onkernel/kernel-go-sdk/commit/9542333108abb522bae00b266c89cc3917884b35))
* update @stainless-api/prism-cli to v5.15.0 ([625476c](https://github.com/onkernel/kernel-go-sdk/commit/625476c16f8298f7fa1c18318f67231906d89a56))

## 0.9.0 (2025-08-07)

Full Changelog: [v0.8.2...v0.9.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.8.2...v0.9.0)

### Features

* **api:** browser instance file i/o ([7308fb8](https://github.com/onkernel/kernel-go-sdk/commit/7308fb8acdcb2d10fb9d09612854a18f18083bcc))
* **client:** support optional json html escaping ([7fb83d1](https://github.com/onkernel/kernel-go-sdk/commit/7fb83d1f154d9502833b4f0a17c9912aa9643c65))

## 0.8.2 (2025-07-23)

Full Changelog: [v0.8.1...v0.8.2](https://github.com/onkernel/kernel-go-sdk/compare/v0.8.1...v0.8.2)

### Features

* **api:** add action name to the response to invoke ([e1bfb1d](https://github.com/onkernel/kernel-go-sdk/commit/e1bfb1d1874a0ad1dbb61c035e3f728b9406df5a))


### Bug Fixes

* **client:** process custom base url ahead of time ([4bef9af](https://github.com/onkernel/kernel-go-sdk/commit/4bef9afe3097d7d07596799efae41ca2acbff038))

## 0.8.1 (2025-07-21)

Full Changelog: [v0.8.0...v0.8.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.8.0...v0.8.1)

### Chores

* **api:** remove deprecated endpoints ([7a659dc](https://github.com/onkernel/kernel-go-sdk/commit/7a659dcfeba2f3881b84934c2ee4780c4fbbb852))

## 0.8.0 (2025-07-16)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.7.1...v0.8.0)

### Features

* **api:** manual updates ([ba0a473](https://github.com/onkernel/kernel-go-sdk/commit/ba0a473799c2c9a00865112a22ae356c57d2c114))
* **client:** expand max streaming buffer size ([41bdefe](https://github.com/onkernel/kernel-go-sdk/commit/41bdefed3d43cb52b431e8255a4a4faa0dfc46f0))


### Chores

* **internal:** fix lint script for tests ([5118fdf](https://github.com/onkernel/kernel-go-sdk/commit/5118fdf7e31a237e8def04c5a41cd0767900deef))
* lint tests in subpackages ([71d9002](https://github.com/onkernel/kernel-go-sdk/commit/71d900220761ff811187d5c03236a0fc5d92fcc6))

## 0.7.1 (2025-07-08)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.7.0...v0.7.1)

### Features

* **api:** manual updates ([2162b0a](https://github.com/onkernel/kernel-go-sdk/commit/2162b0a2396b520acf176b8622b8728e51f8787d))


### Chores

* lint tests ([29d98ac](https://github.com/onkernel/kernel-go-sdk/commit/29d98ac8ec7d9fe704d07ac0bce4b60d0f2ed121))

## 0.7.0 (2025-07-02)

Full Changelog: [v0.6.5...v0.7.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.5...v0.7.0)

### Features

* **api:** headless browsers ([e7f85f9](https://github.com/onkernel/kernel-go-sdk/commit/e7f85f9432b053b513b1581cb2ece0dd78ad3613))

## 0.6.5 (2025-07-02)

Full Changelog: [v0.6.4...v0.6.5](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.4...v0.6.5)

### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([7865ba8](https://github.com/onkernel/kernel-go-sdk/commit/7865ba807c2766ccdd8cb3d59bb8855721b39769))


### Chores

* **ci:** only run for pushes and fork pull requests ([a6b4ae7](https://github.com/onkernel/kernel-go-sdk/commit/a6b4ae749a0e11f8683dd3c354c1e3b3003d738e))

## 0.6.4 (2025-06-27)

Full Changelog: [v0.6.3...v0.6.4](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.3...v0.6.4)

### Features

* **api:** add GET deployments endpoint ([e2c8f14](https://github.com/onkernel/kernel-go-sdk/commit/e2c8f14da32a692adbf18162ceab4b89e1bb2c4f))
* **api:** deployments ([e0dc9c0](https://github.com/onkernel/kernel-go-sdk/commit/e0dc9c0fc621dabb383975656e58a19bbf8c0714))
* **api:** manual updates ([f2ddf2c](https://github.com/onkernel/kernel-go-sdk/commit/f2ddf2c318eb1254e8c15fb5723a92b0ebd9cbb4))

## 0.6.3 (2025-06-25)

Full Changelog: [v0.6.2...v0.6.3](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.2...v0.6.3)

### Features

* **api:** /browsers no longer requires invocation id ([179da50](https://github.com/onkernel/kernel-go-sdk/commit/179da501250b0950198690bf54fabf046264c3f5))

## 0.6.2 (2025-06-24)

Full Changelog: [v0.6.1...v0.6.2](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.1...v0.6.2)

### Features

* **api:** add `since` parameter to deployment logs endpoint ([dc72c81](https://github.com/onkernel/kernel-go-sdk/commit/dc72c81a2e45918d595c6c00843b1a1d0efffdd0))
* **client:** add escape hatch for null slice & maps ([d5e1ad9](https://github.com/onkernel/kernel-go-sdk/commit/d5e1ad9087aecd6b67369b9ebbeb633ad808c129))


### Chores

* fix documentation of null map ([a62b964](https://github.com/onkernel/kernel-go-sdk/commit/a62b9647386501f43e70ad876dc5c0271c4c4709))

## 0.6.1 (2025-06-18)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.6.0...v0.6.1)

### Features

* **api:** add delete_browsers endpoint ([57fb070](https://github.com/onkernel/kernel-go-sdk/commit/57fb070819a948f64af4882595f99fa37f462c70))

## 0.6.0 (2025-06-17)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.5.0...v0.6.0)

### Features

* **api:** update via SDK Studio ([ee8e77a](https://github.com/onkernel/kernel-go-sdk/commit/ee8e77a653cc83084eec067e1060b80e8e99fc27))
* **api:** update via SDK Studio ([1b2efd8](https://github.com/onkernel/kernel-go-sdk/commit/1b2efd8f083758e39ed2bd9bec8d5f6da0642ece))
* **api:** update via SDK Studio ([3094578](https://github.com/onkernel/kernel-go-sdk/commit/3094578598220a1b837274b20e0dd4cb4f36fd2c))
* **api:** update via SDK Studio ([0fec1e5](https://github.com/onkernel/kernel-go-sdk/commit/0fec1e5e8d6b426ade8030ec6836142ddd18cca4))
* **api:** update via SDK Studio ([b44c90e](https://github.com/onkernel/kernel-go-sdk/commit/b44c90ed4698ca9c849bb20aa85579d0fcd36736))
* **api:** update via SDK Studio ([d6e1cd6](https://github.com/onkernel/kernel-go-sdk/commit/d6e1cd618e324beb797495636a79016aa9cfe7b1))
* **api:** update via SDK Studio ([0bc5b00](https://github.com/onkernel/kernel-go-sdk/commit/0bc5b008aad753e6b2a5bf2647b30b389e36c190))
* **api:** update via SDK Studio ([db224f1](https://github.com/onkernel/kernel-go-sdk/commit/db224f133e8935217833037397839b771e901885))
* **api:** update via SDK Studio ([3be7afb](https://github.com/onkernel/kernel-go-sdk/commit/3be7afb4985902b399ed45192a139b8f2a67bcf2))
* **api:** update via SDK Studio ([562f248](https://github.com/onkernel/kernel-go-sdk/commit/562f248c183bc6bd908433c2920b63b596572711))
* **api:** update via SDK Studio ([984cd7f](https://github.com/onkernel/kernel-go-sdk/commit/984cd7f8a9cbbb153e11895e5b8dd5ba79cc5548))
* **api:** update via SDK Studio ([48d0b3d](https://github.com/onkernel/kernel-go-sdk/commit/48d0b3d6938a8e4e9db9148b7dc35d7440e43ae2))
* **api:** update via SDK Studio ([8c6285b](https://github.com/onkernel/kernel-go-sdk/commit/8c6285b9dfff61a9f6b58053026ac18fbde1b66e))
* **client:** add debug log helper ([5c92a71](https://github.com/onkernel/kernel-go-sdk/commit/5c92a71246d10274fa60eb84a0f42121cbca6e7b))


### Chores

* **ci:** enable for pull requests ([72fd885](https://github.com/onkernel/kernel-go-sdk/commit/72fd885ebc2a949148fe7fbc5bc809b716a783fe))

## 0.5.0 (2025-06-04)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.4.0...v0.5.0)

### Features

* **api:** update via SDK Studio ([f0b66f8](https://github.com/onkernel/kernel-go-sdk/commit/f0b66f87541f4bc109f479906c3554bffb38c26f))
* **api:** update via SDK Studio ([d60d333](https://github.com/onkernel/kernel-go-sdk/commit/d60d3337306d90c8c51a54c6abeece2731dd4834))
* **client:** allow overriding unions ([51d877d](https://github.com/onkernel/kernel-go-sdk/commit/51d877d1ae7584b237fdb9238e97949b832c3cc8))


### Bug Fixes

* **client:** cast to raw message when converting to params ([de14358](https://github.com/onkernel/kernel-go-sdk/commit/de14358a78248e642c03bd669f9c361bf6b3c8ba))
* fix error ([7d27985](https://github.com/onkernel/kernel-go-sdk/commit/7d2798511671e285d41943c51c3ac8ebc6a7d6d6))


### Chores

* make go mod tidy continue on error ([91b5021](https://github.com/onkernel/kernel-go-sdk/commit/91b50217a8a08dd06ba64e05d025386d91d586c8))

## 0.4.0 (2025-05-28)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.3.0...v0.4.0)

### Features

* **api:** update via SDK Studio ([ac5cf50](https://github.com/onkernel/kernel-go-sdk/commit/ac5cf50867c042d1da5329f2441855ab89efd686))

## 0.3.0 (2025-05-22)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.2.0...v0.3.0)

### Features

* **api:** update via SDK Studio ([7c22f9e](https://github.com/onkernel/kernel-go-sdk/commit/7c22f9efecdb2548ce58f81454817a33fc249128))
* **api:** update via SDK Studio ([7941c72](https://github.com/onkernel/kernel-go-sdk/commit/7941c7280d3cc6870947d59c96d5e1c55e0ade46))
* **api:** update via SDK Studio ([470093a](https://github.com/onkernel/kernel-go-sdk/commit/470093adf03fc0b077d11f11f13ee802ae4557a6))
* **api:** update via SDK Studio ([5c567c9](https://github.com/onkernel/kernel-go-sdk/commit/5c567c90794bfd2d997685e5102835d65de029e6))


### Chores

* **docs:** grammar improvements ([cb762f8](https://github.com/onkernel/kernel-go-sdk/commit/cb762f85ee2aef349829477128d84e63e41d4449))
* improve devcontainer setup ([8b46076](https://github.com/onkernel/kernel-go-sdk/commit/8b46076b6e1145d9f4ff14e089aa57588f8ef613))

## 0.2.0 (2025-05-21)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0...v0.2.0)

### Features

* **api:** update via SDK Studio ([788354c](https://github.com/onkernel/kernel-go-sdk/commit/788354c6debdf1f4608c31acaa6bf422af3f6306))

## 0.1.0 (2025-05-21)

Full Changelog: [v0.1.0-alpha.6...v0.1.0](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.6...v0.1.0)

### Features

* **api:** update via SDK Studio ([e05f5ec](https://github.com/onkernel/kernel-go-sdk/commit/e05f5ec48a0139acb5f5789124a15916324ceebb))

## 0.1.0-alpha.6 (2025-05-20)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** update via SDK Studio ([8447b94](https://github.com/onkernel/kernel-go-sdk/commit/8447b94117d58db98c50e1e8248121d38c7afd2e))
* **api:** update via SDK Studio ([da5cfff](https://github.com/onkernel/kernel-go-sdk/commit/da5cfff4d62c3dbbac4c2fded65505368097f5a2))

## 0.1.0-alpha.5 (2025-05-20)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update via SDK Studio ([2f5051e](https://github.com/onkernel/kernel-go-sdk/commit/2f5051e328b157c2836d93fcd1e957c315f9d4e7))
* **api:** update via SDK Studio ([368ec0a](https://github.com/onkernel/kernel-go-sdk/commit/368ec0a4e3eb54fd5634327f6234d9b4f0a83a95))

## 0.1.0-alpha.4 (2025-05-20)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** update via SDK Studio ([34dfea0](https://github.com/onkernel/kernel-go-sdk/commit/34dfea07fd8538b3a73f8ee7ea35f3a70e679b3b))


### Bug Fixes

* **client:** correctly set stream key for multipart ([2336290](https://github.com/onkernel/kernel-go-sdk/commit/23362905d1acb3616333407f70bdee259dc66ebe))
* **client:** don't panic on marshal with extra null field ([4208ce7](https://github.com/onkernel/kernel-go-sdk/commit/4208ce7bb382ac9bfb2cefeb406be6302df6cd7a))
* **client:** increase max stream buffer size ([eca7429](https://github.com/onkernel/kernel-go-sdk/commit/eca74291efed0e090b1a7437d832ee12a819f833))
* **client:** use scanner for streaming ([6fad3f9](https://github.com/onkernel/kernel-go-sdk/commit/6fad3f9d43fbb49c718a53ae6736ac273e43b7b3))

## 0.1.0-alpha.3 (2025-05-19)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([4a12204](https://github.com/onkernel/kernel-go-sdk/commit/4a122041257d78d3b583408e912d4d8a5b10cdc0))

## 0.1.0-alpha.2 (2025-05-19)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/onkernel/kernel-go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([72da2c4](https://github.com/onkernel/kernel-go-sdk/commit/72da2c4b53d8f47d6cc0dab3cfbe8707edb989d2))
* **api:** update via SDK Studio ([58da3af](https://github.com/onkernel/kernel-go-sdk/commit/58da3af97eb8b295a54e22e53d8035f00fe09215))
* **api:** update via SDK Studio ([ebab506](https://github.com/onkernel/kernel-go-sdk/commit/ebab506e3b5f8dbe8d30a47b3734c64024341d87))

## 0.1.0-alpha.1 (2025-05-14)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/onkernel/kernel-go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([bb5cfe4](https://github.com/onkernel/kernel-go-sdk/commit/bb5cfe49047afce492f2a21ade2acbe1612e9f3c))


### Chores

* configure new SDK language ([dd0120b](https://github.com/onkernel/kernel-go-sdk/commit/dd0120b0ce673e0fe7c842d39e91f01b8ee8106a))
* update SDK settings ([b1b8645](https://github.com/onkernel/kernel-go-sdk/commit/b1b8645621de70f28e3cf0b3622c1a93159b561e))
