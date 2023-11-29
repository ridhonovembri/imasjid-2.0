import http from '@/api/http'

class Get {
  slides(){
    return http.get('/slides')
  }

  getSlides(){
    return http.get('/getSlides')
  }

  slideById(id) {
    return http.get(`/slideById/${ id }`);
  }

  slideByName(name) {
    // console.log(name)
    return http.get(`/slideByName/${ name }`);
  }

  masjidInfo(){
    return http.get('/masjidInfo')
  }

  masjidInfoById(id) {
    return http.get(`/masjidInfoById/${ id }`);
  }

  masjidConfig(){
    return http.get('/masjidConfig')
  }

  masjidConfigById(id) {
    return http.get(`/masjidConfigById/${ id }`);
  }

  masjidHadist(){
    // console.log('get Hadist')
    return http.get('/masjidHadist')
  }

  masjidHadistRandom(){
    // console.log('get Hadist')
    return http.get('/masjidHadistRandom')
  }

  masjidHadistById(id) {
    return http.get(`/masjidHadistById/${ id }`);
  }

  masjidMarquee(){
    // console.log('get Hadist')
    return http.get('/masjidMarquee')
  }

  masjidMarqueeById(id) {
    return http.get(`/masjidMarqueeById/${ id }`);
  }
  
}

export default new Get();

