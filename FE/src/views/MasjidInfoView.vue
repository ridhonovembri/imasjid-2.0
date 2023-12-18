<template>
  <div class="q-pa-md row items-start q-gutter-md">
    <q-card bordered class="card">
      <q-card-section elevated class="bg-primary text-grey-1">
        <div class="text-h6">Informasi Masjid</div>
      </q-card-section>

      <q-separator />

      <div class="q-gutter-xs" style="margin: 10px">
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.MasjidName"
          label="Nama Masjid"
        />
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.MasjidAddress"
          label="Alamat"
        />
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.City"
          label="Kota"
        />
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.Province"
          label="Propinsi"
        />
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.BkmPic"
          label="BKM Staff"
        />
        <q-input
          standout="bg-teal text-white"
          v-model="masjidInfo.BkmPhone"
          label="Phone"
        />
        <q-btn
          color="white"
          text-color="black"
          label="Simpan"
          @click="updateMasjidInfo"
        />
      </div>
    </q-card>
  </div>
</template>
<script>
import Update from "@/api/http-update";
import Get from "@/api/http-get";
import { Notify } from "quasar";

export default {
  data() {
    return {
      masjidInfo: {
        Id: "",
        MasjidName: "",
        MasjidAddress: "",
        City: "",
        Province: "",
        BkmPic: "",
        BkmPhone: "",
      },
    };
  },
  async mounted() {
    // this.updateMasjidInfo();
    this.getMasjidInfo();
  },
  methods: {
    async updateMasjidInfo() {
      // console.log("masjidInfo", this.masjidInfo);
      // console.log('this.masjidInfo.id', this.masjidInfo.Id)
      // Update.masjidInfo();
      await Update.masjidInfo(this.masjidInfo.Id, this.masjidInfo).then((res) => {

        const status = res.data.status;
        const message = res.data.message;
        // console.log("after update", status);

        if (status) {
          Notify.create({
            message,
            color: "blue-8",
            textColor: "white",
            icon: "cloud_done",
          });
        }
      });
    },

    async getMasjidInfo() {
      this.masjidInfo = (await Get.masjidInfo()).data;
      // console.log('masjidInfo', this.masjidInfo)
    },
  },
};
</script>
<style scoped>
.card {
  width: 100%;
  text-align: left;
  /* max-width: 250px */
}
</style>
