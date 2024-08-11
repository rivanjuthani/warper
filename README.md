# Warper 🌪️
a simple cli to simulate referral installations for 1.1.1.1 WARP+

> Note: Cloudflare will conclude the referral program for WARP+ on November 1, 2024. At that time, new referral codes will no longer be provided. When your device consumes existing referral allotment you will automatically migrate to the WARP free plan. You will not lose connectivity.

## How to use?

Download one of the prebuilt binaries from the releases or build the project from source.

```bash
./warper [referral_link] [amount_of_referrals]
```

Example usage (not a valid referral link):

```bash
./warper https://warp.plus/d24yT 5
```

![13 GB referral credit](https://github.com/rivanjuthani/warper/blob/main/images/warp.jpg?raw=true)

## Build from source

##### Prerequisites:
- Golang 1.22.6 or above
- A valid WARP+ referral link

##### Steps:

First clone the repo & cd into the folder:

```bash
git clone https://github.com/rivanjuthani/warper
cd warper
```

Install required go packages:

```bash
go get
```

Build & run the go binary:

```bash
go build .
./warper
```

## TODO
- Proxy support
- Clean up code to be more readable
- Randomize device data
- Add retry incase adding fails instead of skipping
- Improve output
- Add debug output option