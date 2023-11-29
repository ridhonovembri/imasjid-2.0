<template>
  <router-view/>
</template>

<style>
#app {
  font-family: montserratmedium, Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

@font-face {
    font-family: 'montserratextrabold';
    src: url('./assets/fonts/montserrat/montserrat-extrabold.woff2') format('woff2'),
         url('./assets/fonts/montserrat/montserrat-extrabold.woff') format('woff');
    font-weight: normal;
    font-style: normal;
}

@font-face {
    font-family: 'montserratlight';
    src: url('./assets/fonts/montserrat/montserrat-light.woff2') format('woff2'),
         url('./assets/fonts/montserrat/montserrat-light.woff') format('woff');
    font-weight: normal;
    font-style: normal;
}

@font-face {
    font-family: 'montserratmedium';
    src: url('./assets/fonts/montserrat/montserrat-medium.woff2') format('woff2'),
         url('./assets/fonts/montserrat/montserrat-medium.woff') format('woff');
    font-weight: normal;
    font-style: normal;
}

@font-face {
    font-family: 'montserratregular';
    src: url('./assets/fonts/montserrat/montserrat-regular.woff2') format('woff2'),
         url('./assets/fonts/montserrat/montserrat-regular.woff') format('woff');
    font-weight: normal;
    font-style: normal;
}

/* @font-face {
  font-family: 'Lato-Black';
  src: url('../fonts/lato/Lato-Black.woff2') format('woff2'),
       url('../fonts/lato/Lato-Black.woff') format('woff');
  font-weight: normal;
  font-style: normal;
} */


</style>
