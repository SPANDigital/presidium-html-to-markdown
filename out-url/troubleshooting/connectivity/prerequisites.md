---
title: "Prerequisites¶"
weight: 2
---

To run the sample commands below you will first need to:

- Ensure you have `brew` setup and configured using the [External Link: internal Homebrew @ Apple](https://github.pie.apple.com/homebrew/brew)
- Install [External Link: jq](https://stedolan.github.io/jq/) and [External Link: grpcurl](https://github.com/fullstorydev/grpcurl):


  ```
  brew install jq
  brew install grpcurl

  ```

- Set up [Internal Link: IAS client credentials](../self-service/public_api.html#setting-up-ias-client-environment-variables) and [Internal Link: `kcurl` helper function](../self-service/public_api.html#kcurl)
- Download the Kaffe protobuf source file:


  ```
  curl https://docs.aci.apple.com/kafka/kaffe.proto > kaffe.proto

  ```


