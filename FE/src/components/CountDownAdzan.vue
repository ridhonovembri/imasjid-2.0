<template>
  <div class="text-weight-bold" style="font-size: 54px;">{{ label }} - {{ hour }}.<span style="font-size: 40px;">{{ second }}</span></div>  
</template>
<script>
export default {
  props: ['label', 'hour', 'second'],
  data() {   
    return {    
    };
  },
  async mounted() {
  },
  methods: {    
  },
};
</script>
<style></style>
