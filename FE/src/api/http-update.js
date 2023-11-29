import http from "@/api/http";

class Update {
  slides(id, data) {
    return http.put(`/slide/update/${id}`, data);
  }
  
  masjidInfo(id, data) {
    return http.put(`/masjidInfo/update/${id}`, data);
  }

  masjidConfig(id, data) {
    return http.put(`/masjidConfig/update/${id}`, data);
  }

  masjidHadist(id, data) {
    return http.put(`/masjidHadist/update/${id}`, data);
  }
  
  masjidMarquee(id, data) {
    return http.put(`/masjidMarquee/update/${id}`, data);
  }
}

export default new Update();
