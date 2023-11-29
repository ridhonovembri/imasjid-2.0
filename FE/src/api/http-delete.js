import http from "@/api/http";

class Delete {
  slides(id) {
    return http.delete(`/slide/delete/${id}`);
  }

  masjidInfo(id) {
    return http.delete(`/masjidInfo/delete/${id}`);
  }

  masjidConfig(id) {
    return http.delete(`/masjidConfig/delete/${id}`);
  }

  masjidHadist(id) {
    return http.delete(`/masjidHadist/delete/${id}`);
  }

  masjidMarquee(id) {
    return http.delete(`/masjidMarquee/delete/${id}`);
  }
}

export default new Delete();
