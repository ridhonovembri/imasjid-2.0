<template>
  <!-- <div class="text-h1 text-yellow-5 text-weight-bold">{{ hour }}. {{ second }}</div>   -->
  <div>
    <div class="card-header" >
      <div style="font-size: 60px;">Menanti Adzan {{ prayerName }}</div>
      <div style="font-size: 32px;" v-html="text"></div>
    </div>

    <div class="row">
      <q-card class="card-time">
        <q-item>
          <q-item-section >
            <span style="font-size: 120px;">{{ hour }}</span>
            <span style="font-size: 36px;">Menit</span>
          </q-item-section>

          <q-item-section>
            <span style="font-size: 120px;">{{ second }}</span>
            <span style="font-size: 36px;">Detik</span>
          </q-item-section>
        </q-item>
      </q-card>
    </div>
  </div>
</template>
<script>
export default {
  props: ["hour", "second", "prayerName", "text"],
  data() {
    return {};
  },
  async mounted() {},
  methods: {},
};
</script>
<style scoped>
.card-header {
  /* width: 100%; */
  /* border: 1px solid red; */
  /* color: #ffd152; */
  max-width: 100%;
  height: 27vh;
  font-size: 32px;
  font-family: "montserratmedium";
  text-align: center;
  /* letter-spacing: 10px; */
  /* box-shadow: 0 0 29px 0 rgba(18, 66, 101, 0.5); */
  /* background: rgba(168, 218, 220, 0.3); */
  /* background-color: rgba(0, 0, 0, 0.2); */
  color: #fff1e6;
  /* color: #F1FAEE; */
  background: rgba(1, 51, 122, 0.3);
}

.card-time {
  width: 100%;
  font-family: "montserratextrabold";
  text-align: center;
  letter-spacing: 10px;
  font-size: 60px;
  margin: auto;
  /* background-color: rgba(0, 0, 0, 0.2); */
  /* color: #fff1e6; */
  /* color:rgb(0,128,128); */
  /* color: #ffd152; */
  /* background-color: none; */
  /* box-shadow: 0 0 29px 0 rgba(18, 66, 101, 0.5); */
  background: rgba(1, 51, 122, 0.3);
  color: #fff1e6;
  /* background: #124265;  */

  
}

/* .prayer-name {
  color: #ffd152;
  margin: 0;
  padding: 0;
  line-height: 1;
  font-family: "montserratmedium";
  font-weight: 700;
  letter-spacing: 3px;
} */
/* .card {
  box-shadow: 0 0 29px 0 rgba(18, 66, 101, 0.5);
  background: rgba(168, 218, 220, 0.3);
} */

/* .card-header {
  background: #1d3557;
  color: #f1faee;
  font-size: 60px;
  font-family: "montserratmedium";
  text-align: center;
  letter-spacing: 10px;
} */

/* .card-title {
  font-size: 270px;
  font-weight: 700;
  font-family: "montserratbold";
  color: #495057;
} */
</style>
