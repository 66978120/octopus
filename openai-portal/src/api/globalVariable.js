let DOMAIN
if (process.env.NODE_ENV === 'development') { DOMAIN = 'http://192.168.202.73' }
else { DOMAIN = window.location.protocol + '//' + document.domain }
export default {
  DOMAIN
}