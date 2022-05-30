<template>
  <v-container>
    <h1>Add Text Proposal</h1>

    <form @submit.prevent="onSubmitProposal()">
      <v-text-field label="Title" placeholder="Title" v-model="title" :rules="[nonemptyTitle]" />

      <v-text-field
        label="Initial Deposit"
        placeholder="Initial Deposit"
        v-model="initialDeposit"
        :rules="[money]"
      />

      <v-textarea
        label="Description"
        placeholder="Description"
        v-model="description"
        hint="Markdown is supported"
        :rules="[nonemptyDescription]"
      />

      <v-row justify="center">
        <v-btn
          color="primary"
          type="submit"
          :disabled="loading"
          :loading="loading"
        >
          Submit
        </v-btn>
      </v-row>
    </form>

    <CustomAlert :loading="loading" :result="result" :error="error" />
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { submitProposal } from "@/utils/gov";
import { BigNumber } from "ethers";
import CustomAlert from "@/components/CustomAlert.vue";
import { parseUnits } from "ethers/lib/utils";

@Component({
  name: "TextProposal",
  components: {
    CustomAlert,
  },
})
export default class TextProposal extends Vue {
  loading = false;
  result: any = null;
  error: string | null = null;

  title = "";
  description = "";
  initialDeposit = "1";

  parseAmount(): BigNumber {
    const decimals = this.$store.state.config.tft_decimals || 0;
    const amountBN = parseUnits(this.initialDeposit || "0", decimals);
    if (amountBN.lte(0)) {
      throw new Error("amount must be positive")
    }
    return amountBN
  }

  money() {
    try {
      this.parseAmount()
    } catch (err: any) {
      return err.message
    }
  }

  nonemptyTitle() {
    if (this.title == "") {
      return "title can't be blank"
    }
    return true
  }

  nonemptyDescription() {
    if (this.description == "") {
      return "description can't be blank"
    }
    return true
  }

  onSubmitProposal() {
    this.loading = true;
    this.result = null;
    this.error = null;

    try {
      const { title, description, initialDeposit } = this;
      submitProposal(
        this.$store.state.config.tendermint_rpc,
        this.$store.state.config.cosmos_rest,
        this.$store.state.config.gas_price,
        this.$store.state.config.chain_id,
        { title, description },
        parseUnits(initialDeposit, this.$store.state.config.tft_decimals),
        this.$store.state.config.proposal_denom
      )
        .then((res) => {
          this.result = "Proposal added succefully!";
        })
        .catch((err) => {
          this.error = err.message;
        })
        .finally(() => {
          this.loading = false;
        });
    } catch (err: any) {
      this.loading = false;
      this.error = err.message;
    }
  }
}
</script>