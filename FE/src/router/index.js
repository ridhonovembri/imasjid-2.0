import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import SlidesView from "@/views/SlidesView.vue";
import MasjidInfoView from "@/views/MasjidInfoView.vue";
import MasjidConfigView from '@/views/MasjidConfigView.vue'
import HadistView from "@/views/HadistView.vue";
import DashboardView from "@/views/DashboardView.vue";
import MarqueeView from "@/views/MarqueeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: DashboardView,
    children: [
      {
        path: "/slides",
        name: "slides",
        component: SlidesView,
      },
      {
        path: "/masjidinfo",
        name: "masjidinfo",
        component: MasjidInfoView,
      },
      {
        path: "/masjidconfig",
        name: "masjidconfig",
        component: MasjidConfigView,
      },
      {
        path: "/hadist",
        name: "hadist",
        component: HadistView,
      },
      {
        path: "/marquee",
        name: "marquee",
        component: MarqueeView,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
