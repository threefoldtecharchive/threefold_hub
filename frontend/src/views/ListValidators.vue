<template>
  <v-container>
    <h1>Validators</h1>
    <v-alert
      type="error"
      v-if="error != null"
    >
    {{ error }}
    </v-alert>
    <v-data-table :headers="headers" :items="validators" :loading="loading">

      <template v-slot:[`item.commission`]="{ item }">
        {{ +item.commission.commission_rates.rate * 100 }}%
      </template>

      <template v-slot:[`item.details`]="{ item }">
        <v-btn
          color="primary"
          @click="$router.push('/delegate/' + item.operator_address)"
        >
          Delegate
        </v-btn>
      </template>
    </v-data-table>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import {
  CosmosStakingV1Beta1QueryValidatorsResponse,
} from "@/rest/cosmos";
import { listValidators, parameters, tally } from "@/utils/gov";
import { formatUnits } from "ethers/lib/utils";
import { parseUnits } from "@/utils/money";
import VoteCircle from "@/components/VoteCircle.vue";

@Component({
  name: "ListGov",
  components: {
    VoteCircle,
  },
})
export default class ListGov extends Vue {
  headers: { text: string; value: string }[] = [
    { text: "Address", value: "operator_address" },
    { text: "Voting power", value: "tokens" },
    { text: "Commission", value: "commission" },
    { text: "Delegate", value: "details" },
  ];

  validators: CosmosStakingV1Beta1QueryValidatorsResponse["validators"] = [];
  loading = false;
  error: string | null = null;
  
  async normalizeStakedTokens() {
    let validators = this.validators || [];
    for (const validator of validators) {
      validator.tokens = formatUnits(validator.tokens || "0", this.$store.state.config.tft_decimals) + " " + this.$store.state.config.tft_denom
    }
  }

  created() {
    this.loading = true;
    listValidators(this.$store.state.config.cosmos_rest)
      .then((res: CosmosStakingV1Beta1QueryValidatorsResponse) => {
        this.validators = res.validators;
        this.normalizeStakedTokens()
      })
      .catch((err) => {
        this.error = "Failed to list validators (refresh to try again): " + err.message
      })
      .finally(() => {
        this.loading = false;
      });
  }
}
</script>
