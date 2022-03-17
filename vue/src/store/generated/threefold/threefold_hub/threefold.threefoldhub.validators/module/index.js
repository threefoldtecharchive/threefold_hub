// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateValidator } from "./types/validators/tx";
import { MsgCreateValidator } from "./types/validators/tx";
import { MsgDeleteValidator } from "./types/validators/tx";
const types = [
    ["/threefold.threefoldhub.validators.MsgUpdateValidator", MsgUpdateValidator],
    ["/threefold.threefoldhub.validators.MsgCreateValidator", MsgCreateValidator],
    ["/threefold.threefoldhub.validators.MsgDeleteValidator", MsgDeleteValidator],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateValidator: (data) => ({ typeUrl: "/threefold.threefoldhub.validators.MsgUpdateValidator", value: MsgUpdateValidator.fromPartial(data) }),
        msgCreateValidator: (data) => ({ typeUrl: "/threefold.threefoldhub.validators.MsgCreateValidator", value: MsgCreateValidator.fromPartial(data) }),
        msgDeleteValidator: (data) => ({ typeUrl: "/threefold.threefoldhub.validators.MsgDeleteValidator", value: MsgDeleteValidator.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
