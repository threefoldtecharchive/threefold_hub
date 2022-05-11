// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgExecuteTransaction } from "./types/multisigwallet/tx";
import { MsgAddMember } from "./types/multisigwallet/tx";
import { MsgUpdateMinSigns } from "./types/multisigwallet/tx";
import { MsgSignMemberTransaction } from "./types/multisigwallet/tx";
import { MsgSignTransaction } from "./types/multisigwallet/tx";
import { MsgRemoveMember } from "./types/multisigwallet/tx";
import { MsgCreateWallet } from "./types/multisigwallet/tx";
import { MsgCreateTransaction } from "./types/multisigwallet/tx";


const types = [
  ["/threefoldtech.threefoldhub.multisigwallet.MsgExecuteTransaction", MsgExecuteTransaction],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgAddMember", MsgAddMember],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgUpdateMinSigns", MsgUpdateMinSigns],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgSignMemberTransaction", MsgSignMemberTransaction],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgSignTransaction", MsgSignTransaction],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgRemoveMember", MsgRemoveMember],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgCreateWallet", MsgCreateWallet],
  ["/threefoldtech.threefoldhub.multisigwallet.MsgCreateTransaction", MsgCreateTransaction],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgExecuteTransaction: (data: MsgExecuteTransaction): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgExecuteTransaction", value: MsgExecuteTransaction.fromPartial( data ) }),
    msgAddMember: (data: MsgAddMember): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgAddMember", value: MsgAddMember.fromPartial( data ) }),
    msgUpdateMinSigns: (data: MsgUpdateMinSigns): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgUpdateMinSigns", value: MsgUpdateMinSigns.fromPartial( data ) }),
    msgSignMemberTransaction: (data: MsgSignMemberTransaction): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgSignMemberTransaction", value: MsgSignMemberTransaction.fromPartial( data ) }),
    msgSignTransaction: (data: MsgSignTransaction): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgSignTransaction", value: MsgSignTransaction.fromPartial( data ) }),
    msgRemoveMember: (data: MsgRemoveMember): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgRemoveMember", value: MsgRemoveMember.fromPartial( data ) }),
    msgCreateWallet: (data: MsgCreateWallet): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgCreateWallet", value: MsgCreateWallet.fromPartial( data ) }),
    msgCreateTransaction: (data: MsgCreateTransaction): EncodeObject => ({ typeUrl: "/threefoldtech.threefoldhub.multisigwallet.MsgCreateTransaction", value: MsgCreateTransaction.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
