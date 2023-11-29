<template>
  <div class="text-grey-1 block" v-for="hadist in hadists" :key="hadist.Id">
    <!-- <h3>{{ hadist.HadistContent }}</h3>
    <h4>{{ hadist.HadistFrom }}</h4> -->
    <div style="font-size: 48px;">{{ hadist.HadistContent }}</div>
    <div style="font-size: 36px;">{{ hadist.HadistFrom }}</div>
  </div>
</template>
<script>
import Get from "@/api/http-get";
export default {  
  data() {
    return {
      hadists: [
        {
          HadistContent: "",
          HadistFrom: "",
        },
      ],
      interval: 0,
    };
  },
  setup(props) {
    // setup() receives props as the first argument.
    // console.log("interval hadist", props.interval);
  },
  async mounted() {

    await this.getMasjidConfig();

    setInterval(() => this.getHadist(), this.interval);
  },
  methods: {  
    async getMasjidConfig(){
      const result = (await Get.masjidConfig()).data.IntervalHadist;
      // console.log('result', result)

      this.interval = result
    },

    async getHadist() {
      // console.log('show hadist')
      this.hadists = (await Get.masjidHadistRandom()).data;
      // console.log('hadists', this.hadists)
    },
  },
};
</script>
<style scoped>
.block {
  /* border: 1px solid black; */
  /* padding: 0px; */
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 100%;
  height: 450px;
  /* padding: 0px; */
  background-color: rgba(0, 0, 0, 0.2);
  /* background-color: black; */
  /* text-align: center; */
}
</style>
