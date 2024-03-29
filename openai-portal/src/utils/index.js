/**
 * Created by PanJiaChen on 16/11/18.
 */

/**
 * Parse the time to string
 * @param {(Object|string|number)} time
 * @param {string} cFormat
 * @returns {string | null}
 */
export function parseTime(time, cFormat) {
  if (arguments.length === 0 || !time) {
    return null
  }
  const format = cFormat || '{y}-{m}-{d} {h}:{i}:{s}'
  let date
  if (typeof time === 'object') {
    date = time
  } else {
    if ((typeof time === 'string')) {
      if ((/^[0-9]+$/.test(time))) {
        // support "1548221490638"
        time = parseInt(time)
      } else {
        // support safari
        // https://stackoverflow.com/questions/4310953/invalid-date-in-safari
        time = time.replace(new RegExp(/-/gm), '/')
      }
    }

    if ((typeof time === 'number') && (time.toString().length === 10)) {
      time = time * 1000
    }
    date = new Date(time)
  }
  const formatObj = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay()
  }
  const time_str = format.replace(/{([ymdhisa])+}/g, (result, key) => {
    const value = formatObj[key]
    // Note: getDay() returns 0 on Sunday
    if (key === 'a') { return ['日', '一', '二', '三', '四', '五', '六'][value] }
    return value.toString().padStart(2, '0')
  })
  return time_str
}

/**
 * @param {number} time
 * @param {string} option
 * @returns {string}
 */
export function formatTime(time, option) {
  if (('' + time).length === 10) {
    time = parseInt(time) * 1000
  } else {
    time = +time
  }
  const d = new Date(time)
  const now = Date.now()

  const diff = (now - d) / 1000

  if (diff < 30) {
    return '刚刚'
  } else if (diff < 3600) {
    // less 1 hour
    return Math.ceil(diff / 60) + '分钟前'
  } else if (diff < 3600 * 24) {
    return Math.ceil(diff / 3600) + '小时前'
  } else if (diff < 3600 * 24 * 2) {
    return '1天前'
  }
  if (option) {
    return parseTime(time, option)
  } else {
    return (
      d.getMonth() +
      1 +
      '月' +
      d.getDate() +
      '日' +
      d.getHours() +
      '时' +
      d.getMinutes() +
      '分'
    )
  }
}

/**
 * @param {string} url
 * @returns {Object}
 */
export function param2Obj(url) {
  const search = decodeURIComponent(url.split('?')[1]).replace(/\+/g, ' ')
  if (!search) {
    return {}
  }
  const obj = {}
  const searchArr = search.split('&')
  searchArr.forEach(v => {
    const index = v.indexOf('=')
    if (index !== -1) {
      const name = v.substring(0, index)
      const val = v.substring(index + 1, v.length)
      obj[name] = val
    }
  })
  return obj
}
export function formatDuring(mss) {
  mss = mss * 1000
  var days = parseInt(mss / (1000 * 60 * 60 * 24))
  days = days === 0 ? '' : days + 'd'
  var hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  hours = hours === 0 ? '' : hours + 'h'
  var minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60))
  minutes = minutes === 0 ? '' : minutes + 'm'
  var seconds = Math.round((mss % (1000 * 60)) / 1000) + 's'
  seconds = seconds === 0 ? '' : seconds
  return days + hours + minutes + seconds
}
// 获取浏览器url中指定字段
export function GetUrlParam(paraName) {
  let query = location.href
  query = query.replace('?token', '&token')
  var vars = query.split("&");
  for (var i = 0; i < vars.length; i++) {
    var pair = vars[i].split("=");
    if (pair[0] === paraName) {
      pair[1] = decodeURIComponent(pair[1])
      return pair[1];
    }
  }
  return null;
}
// 刷新页面清空进度条
export function clearProgress(paraName) {
  var storage = window.sessionStorage;
  for (var i = 0, len = storage.length; i < len; i++) {
    var key = storage.key(i);
    if (key !== 'space' && key !== 'platform') {
      storage.removeItem(key)
    }
  }
}

// 生成随机名字
export function randomName(val) {
  var myDate = new Date();
  const year = myDate.getFullYear().toString()
  let month = (myDate.getMonth() + 1).toString()
  if (month < 10) {
    month = '0' + (myDate.getMonth() + 1).toString()
  }
  let date = myDate.getDate().toString()
  if (date < 10) {
    date = '0' + myDate.getDate().toString()
  }
  var charactors = "ab1cd2ef3gh4ij5kl6mn7opq8rst9uvw0xyz";
  var value = ''; var i;
  for (let j = 1; j <= 4; j++) {
    i = parseInt(35 * Math.random());
    value = value + charactors.charAt(i);
  }
  const name = val + '-' + year + month + date + '-' + value
  return name
}
