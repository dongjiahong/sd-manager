export default {
  // 计算两天时间相差天数
  dateDifference(endDate) {
    if (endDate == null || endDate == "") {
      return "0";
    }
    let eDate = Date.parse(endDate.replace(/-/g, "/"));
    let curDate = new Date();
    let diff = eDate - curDate;

    if (diff > 0) {
      return Math.floor(diff / (24 * 3600 * 1000)).toString();
    }
    return "0";
  },
  deepCopy(obj) {
    var result = Array.isArray(obj) ? [] : {};
    for (var key in obj) {
      if (obj.hasOwnProperty(key)) {
        if (typeof obj[key] === "object" && obj[key] !== null) {
          result[key] = this.deepCopy(obj[key]); //递归复制
        } else {
          result[key] = obj[key];
        }
      }
    }
    return result;
  },
  formatDate(date) {
    if (date == "") {
      return "";
    }
    let d = new Date(date);
    let year = d.getFullYear();
    let month = d.getMonth() + 1;
    let day = d.getDate();

    let format =
      year +
      "-" +
      (month <= 9 ? "0" + month : month) +
      "-" +
      (day <= 9 ? "0" + day : day);
    return format;
  }
};
