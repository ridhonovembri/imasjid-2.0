<template>
  <div class="q-pa-md">
    <q-layout view="hHh Lpr lff">
      <q-header elevated class="bg-primary">
        <q-toolbar>
          <q-btn
            flat
            round
            dense
            icon="menu"
            class="q-mr-sm"
            @click="drawer = !drawer"
          />
          <q-toolbar-title style="text-align: left"
            >Masjid Dashboard</q-toolbar-title
          >
        </q-toolbar>
      </q-header>

      <q-drawer
        v-model="drawer"
        show-if-above
        :width="200"
        :breakpoint="500"        
      >
        <q-scroll-area class="fit">
          <q-list>
            <template v-for="(menuItem, index) in menuList" :key="index">
              <q-item clickable :active="menuItem.label === 'Outbox'" v-ripple :to="menuItem.to">
                <q-item-section avatar>
                  <q-icon :name="menuItem.icon" />
                </q-item-section>
                <q-item-section style="text-align: left;">
                  {{ menuItem.label }}
                </q-item-section>
              </q-item>
              <q-separator :key="'sep' + index" v-if="menuItem.separator" />
            </template>
          </q-list>
        </q-scroll-area>
      </q-drawer>

      <q-page-container>
        <q-page padding>
          <router-view />
        </q-page>
      </q-page-container>
    </q-layout>
  </div>
</template>
<script>
export default {
  data() {
    return {
      drawer: false,
      menuList: [
        {
          icon: "inbox",
          label: "Informasi",
          to: "/masjidinfo",
          separator: false,
        },
        {
          icon: "settings",
          label: "Konfigurasi",
          to: "/masjidconfig",
          separator: true,
        },
        {
          icon: "send",
          label: "Pengumuman",
          to: '/marquee',
          separator: false,
        },
        {
          icon: "menu_book",
          label: "Hadist Pilihan",
          to: '/hadist',
          separator: false,
        },
        {
          icon: "slideshow",
          label: "Slides",
          to: '/slides', 
          separator: false,
        },

      ],
    };
  },
};
</script>
<style></style>
