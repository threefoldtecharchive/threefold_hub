<template>
  <v-container>
    <h1>BSC transactions</h1>
    <v-alert
      type="error"
      v-if="permError != null"
    >
    {{ permError }}
    </v-alert>
    <v-tabs v-model="tab">
      <v-tab v-for="t in tabs" :key="t.symbol">
        {{ t.label }}
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model="tab">
      <v-tab-item v-for="t in tabs" :key="t.symbol">
        <v-data-table :headers="t.headers" :items="list[t.symbol]">
          <template v-slot:[`item.erc20_token.amount`]="{ item }">
            {{ normalize(item.erc20_token.amount) }}
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <v-tooltip top>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  fab
                  small
                  dark
                  color="red"
                  v-bind="attrs"
                  v-on="on"
                  @click="onCancelSendToEth(item.id)"
                >
                  <v-icon> mdi-delete </v-icon>
                </v-btn>
              </template>

              <span> Cancel SendToEth Transaction </span>
            </v-tooltip>
          </template>
        </v-data-table>
      </v-tab-item>
    </v-tabs-items>
    <v-snackbar
      :value="error != null"
      color="red"
      absolute
      left
      top
    >
      {{ error }}
      <template v-slot:action="{ attrs }">
        <v-btn
          color="yellow"
          text
          v-bind="attrs"
          @click="error = null"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>

  </v-container>
</template>

<script lang="ts">
import { GravityV1QueryPendingSendToEthResponse } from "@/rest/cosmos";
import { BigNumber } from "ethers";
import { Component, Vue } from "vue-property-decorator";
import { formatUnits } from "ethers/lib/utils";
import { pendingSendToEth, cancelSendToEth } from "../utils";
import { runInThisContext } from "vm";

@Component({
  name: "ListEth",
})
export default class ListEth extends Vue {
  readonly tabs: {
    label: string;
    symbol: keyof GravityV1QueryPendingSendToEthResponse;
    headers: { text: string; value: string }[];
  }[] = [
    {
      label: "Batched",
      symbol: "transfers_in_batches",
      headers: [
        { text: "ID", value: "id" },
        { text: "Destination", value: "dest_address" },
        { text: "Amount", value: "erc20_token.amount" },
      ],
    },
    {
      label: "Unbatched",
      symbol: "unbatched_transfers",
      headers: [
        { text: "ID", value: "id" },
        { text: "Destination", value: "dest_address" },
        { text: "Amount", value: "erc20_token.amount" },
        { text: "Actions", value: "actions" },
      ],
    },
  ];

  list: GravityV1QueryPendingSendToEthResponse = {};
  tab: string | null = null;
  error: string | null = null;
  permError: string | null = null;
  created() {
    // sometimes created is executed before keplr gets injected
    this.updateWhenLoaded();
  }
  normalize(tokens: string): string {
    return formatUnits(BigNumber.from(tokens), this.$store.state.config.tft_decimals);
  }

  updateWhenLoaded() {
    if (document.readyState === "complete") {
      this.updateList();
    } else {
      setTimeout(this.updateWhenLoaded.bind(this), 500);
    }
  }
  updateList() {
    this.permError = null;
    pendingSendToEth(this.$store.state.config.cosmos_rest, this.$store.state.config.chain_id)
      .then((res: GravityV1QueryPendingSendToEthResponse) => {
        this.list = res;
      })
      .catch((err) => {
        this.permError = "Couldn't list BSC transcations (refresh to try again): " + err.message
      });
  }

  onCancelSendToEth(txId: string) {
    const cancel = confirm(
      `Are you sure you want to cancel transaction(${txId})?`
    );

    if (!cancel) return;

    this.error = null
    cancelSendToEth(
      this.$store.state.config.tendermint_rpc,
      this.$store.state.config.cosmos_rest,
      this.$store.state.config.gas_price,
      this.$store.state.config.chain_id,
      txId
      )
      .then((_) => {
        this.updateList();
      })
      .catch((err) => {
        this.error = "Couldn't cancel the BSC transaction: " + err.message
      });
  }
}
</script>
