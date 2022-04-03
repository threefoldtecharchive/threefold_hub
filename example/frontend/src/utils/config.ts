
import { BigNumber } from "ethers";
import { parseUnits } from "ethers/lib/utils";

export interface Config {
    gravity_contract_address: string;
    tft_token_contract_address: string;
    bridge_fees: BigNumber;
    tft_decimals: number;
    tft_denom: string;
    
    cosmos_rest: string;
    tendermint_rpc: string;
}

async function loadJSONConfig(): Promise<{[key: string]: any}> {
    const runtimeConfig = await fetch('/config.json');
    return await runtimeConfig.json();
}

async function validateConfig(config: {[key: string]: any}) {
    const props = [
        "BRIDGE_FEES",
        "TFT_TOKEN_CONTRACT_ADDRESS",
        "GRAVITY_CONTRACT_ADDRESS",
        "TFT_DECIMALS",
        "TFT_DENOM",
        "COSMOS_REST",
        "TENDERMINT_RPC"
    ]
    const numbers = ["BRIDGE_FEES", "TFT_DECIMALS"]
    console.log(config)
    for (const prop of props) {
        if (config[prop] === undefined) {
            throw new Error(prop + " is required and not present in the env vars")
        }
    }
    for (const prop of numbers) {
        if (isNaN(+(config[prop]))) {
            throw new Error(prop + "=" + config[prop] as string + " is not a valid number")
        }
    }
}

export async function loadConfig(): Promise<Config> {
    const config = await loadJSONConfig();
    validateConfig(config);
    const tft_decimals = +(config["TFT_DECIMALS"] as string);
    return {
        bridge_fees: parseUnits(config["BRIDGE_FEES"], tft_decimals),
        tft_token_contract_address: config["TFT_TOKEN_CONTRACT_ADDRESS"],
        gravity_contract_address: config["GRAVITY_CONTRACT_ADDRESS"],
        tft_decimals: tft_decimals,
        tft_denom: config["TFT_DENOM"],
        cosmos_rest: config["COSMOS_REST"] as string,
        tendermint_rpc: config["TENDERMINT_RPC"] as string,
    }
}