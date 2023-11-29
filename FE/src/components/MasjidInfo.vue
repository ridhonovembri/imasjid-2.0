<template>
  <div class="absolute-top-left q-pa-xs text-left">
    <div class="name">
      {{ MasjidName }}
    </div>
    <div class="address">
      {{ MasjidAddress }} - {{ City }} - {{ Province }}
    </div>
    <div class="date">
      {{ Today }} | {{ TodayHijrah }}
    </div>
  </div>
</template>
<script>
export default {
  props: [
    "MasjidName",
    "MasjidAddress",
    "City",
    "Province",
    "Today",
    "TodayHijrah",
  ],
  data() {
    return {};
  },
  async mounted() {},
  methods: {},
};
</script>

<style scoped>

.name {
  font-size: 40px;
  color: #ffd152;
  margin: 0;
  padding: 0;
  line-height: 1;
  font-family: 'montserratextrabold';
  font-weight: 700;
  letter-spacing: 3px;
}

.address {
  font-size: 20px;
  padding-top: 5px;  
  color: #fff1e6;
  letter-spacing: 2px;
  font-family: 'montserratmedium';
  font-weight: 700;
}

.date {
  font-size: 18px;
  padding-top: 2px;  
  color: #fff1e6;
  letter-spacing: 2px;
  font-family: 'montserratmedium';
  font-weight: 700;
}

</style>
