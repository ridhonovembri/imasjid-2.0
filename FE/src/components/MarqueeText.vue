<template>
  <div class="text-grey-1" :style="{'font-size': fontSize}">
    <marquee>{{ text }} </marquee>
  </div>
</template>

<script>
import Get from '@/api/http-get'

export default {
  props: ["fontSize"],
  data() {
    return {
      text:'',
      fontSize: this.fontSize>0? this.fontSize + 'px': '36px'
    };
  },
  async mounted() {
    this.getText()
  },
  methods: {
    async getText(){
      this.marquees = (await (Get.masjidMarquee())).data
      // console.log('this.marquee', this.marquees)

      let marqueeText = this.marquees.map(function(item){
        return item.MarqueeText
      })

      this.text = marqueeText.join(`${'\xa0'.repeat(25)}`)
      // console.log('marqueeText', marqueeText)
    }
  },
};
</script>
<style scoped>
</style>
