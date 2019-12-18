import axios from 'axios';

axios.defaults.headers.post['Content-Type'] = 'application/x-www-from-urlencode'
axios.defaults.timeout = 5000

// add a request interceptor
axios.interceptors.request.use((config)=>{
  return config
},(error) => {
  return Promise.reject(error);
})

// add a response interceptor
axios.interceptors.response.use((response) => {
  if (!response.status === 200) {
    return Promise.reject(response)
  }
  return response
}, (error) => {
  return Promise.reject(error);
})

export default async (url = '', params = {}, method = 'get') => {
  method = method.toLowerCase()
  if (method === 'get') {
    let paramArr = [] 
    for (let [key, value] of Object.entries(params)) {
      paramArr.push(key + '=' + value)
    }
    if (paramArr.length > 0) {
      url += '?' + paramArr.join('&')
    }

    return new Promise((resolve, reject) => {
      axios.get(url).then(response => {
        resolve(response.data)
      }, err=>{
        reject(err)
      }).catch((error) => {
        reject(error)
      })
    })
  } else if (method === 'post') {
    return new Promise((resolve, reject) => {
      axios.post(url, params).then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
    })
  } else {
    return Promise.reject('传递参数错误')
  }
}