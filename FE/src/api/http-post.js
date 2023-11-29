import http from "@/api/http";

class Post {
  slides(data) {
    // console.log('api', data)
    return http.post("/slide/post", data);
  }

  masjidInfo(data) {
    // console.log('api', data)
    return http.post("/masjidInfo/post", data);
  }

  masjidConfig(data) {
    // console.log('api', data)
    return http.post("/masjidConfig/post", data);
  }

  masjidHadist(data) {
    // console.log('api', data)
    return http.post("/masjidHadist/post", data);
  }

  masjidMarquee(data) {
    // console.log('api', data)
    return http.post("/masjidMarquee/post", data);
  }

}

export default new Post();
