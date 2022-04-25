import { message, notification } from 'antd'
import Axios, { AxiosInstance, AxiosRequestTransformer } from 'axios'
import { createBrowserHistory } from 'history'
import { isNil } from 'lodash'
import qs from 'qs'
import { createContext, useContext } from 'react'
import { useMutation, useQuery } from 'react-query'
import { JsonToSnake, superFilter } from './helper'

import Storage from './storage'

enum ResultCode {
  SUCCESS,
  ERROR = 7,
  SYSTEM_ERROR = 100000,
  PARAMS_ERROR
}

const history = createBrowserHistory()
// const navigate = useNavigate()
// console.log('baseurl:', import.meta.env.VITE_BASE_URL)
const axios = Axios.create({
  baseURL: String(import.meta.env.VITE_BASE_URL),
  timeout: 5000,
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
  }
})

axios.interceptors.request.use((config) => {
  const token = Storage.get('accessToken', false)
  if (token) {
    // config.headers = { Authorization: `Bearer ${token}` }
    config.headers['x-token'] = token
  }
  // 每次请求带上时间戳 防刷处理
  // if (config.method === 'get') {
  //     config.params = {
  //         ...config.params,
  //         // @ts-ignore
  //         timestamp: Date.parse(new Date()) / 1000
  //     };
  // }
  return config
})

// response interceptor
axios.interceptors.response.use(
  (response: { data: any; status: number; statusText: any }) => {
    const data = response.data
    // console.log('response:', response)
    if (response.status === 200) {
      switch (data.code) {
        case ResultCode.SUCCESS:
          return data?.data

        case ResultCode.PARAMS_ERROR:
          if (typeof data?.data == 'string') {
            message.warning(data.data)
          }
          if (typeof data?.data == 'object') {
            const values: any[] = Object.values(data?.data)
            values.forEach((value) => {
              message.warning(value)
            })
          }
          return null

        default:
          // eslint-disable-next-line no-case-declarations
          const resultData = data?.data
          // eslint-disable-next-line no-case-declarations
          let resultMsg: string = data.msg
          if (resultData && Object.keys(resultData).length > 0) {
            resultMsg = JSON.stringify(resultData)
          }
          message.warning(resultMsg)
          return null
      }
    }
    // console.log('statusText', response.statusText)
    // notification.error({
    //   message: `请求错误 ${response.statusText}: ${response}`,
    //   description: data || response.statusText || 'Error'
    // })

    return Promise.reject(new Error(response.statusText || 'Error'))
  },
  (error: { response: { status: any; data: { msg: any } } }) => {
    // console.log('err:', error, error.response) // for debug

    if (error.response?.status) {
      switch (error.response.status) {
        // 401: 未登录
        // 未登录则跳转登录页面，并携带当前页面的路径
        // 在登录成功后返回当前页面，这一步需要在登录页操作。
        case 401:
          Storage.remove('accessToken')
          history.push('/login')
          break
        // 403 token过期
        // 登录过期对用户进行提示
        // 清除本地token和清空vuex中token对象
        // 跳转登录页面
        case 403:
          Storage.remove('accessToken')
          history.push('/login')
          break
        // 404请求不存在
        case 404:
          notification.error({
            message: `请求不存在`,
            description: error.response.data?.msg || 'Error'
          })
          break
        case 406:
          notification.error({
            message: `请求参数有误`,
            description: error.response.data?.msg || 'Error'
          })
          break
        default:
          notification.error({
            message: `请求错误`,
            description: error.response.data?.msg || 'Error'
          })
      }
    }

    // throw new Error(error);
    // return Promise.reject(error)
    return null
  }
)

export const AxiosContext = createContext<AxiosInstance>(
  new Proxy(axios, {
    apply: () => {
      throw new Error('You must wrap your component in an AxiosProvider')
    },
    get: () => {
      throw new Error('You must wrap your component in an AxiosProvider')
    }
  })
)

export const useAxios = () => {
  return useContext(AxiosContext)
}

const transformPagination = (pagination: any) => {
  if (!pagination) return

  const current = pagination.current ? pagination.current : pagination.defaultCurrent
  const pageSize = pagination.pageSize ? pagination.pageSize : pagination.defaultPageSize

  // let offset = 0
  // if (current && pageSize) {
  //   offset = (current - 1) * pageSize
  // }

  return {
    page: current,
    pageSize
  }
}

const transformFilters = (filters: any) => {
  if (!filters) return
  let result: any[] = []
  for (const key in filters) {
    if (!filters[key] || filters[key] === null) continue
    result = [...result, [key + ':eq:' + filters[key]]]
  }
  return result
}

const transformSorter = (sorter: any) => {
  if (!sorter) return

  let result = ''
  if (sorter.field && sorter.order) {
    let order = 'desc'
    if (sorter.order === 'ascend') order = 'asc'
    result = sorter.field + ' ' + order
  }

  return result
}

const transformParams = (obj: any) => {
  if (!obj) return {}
  return JsonToSnake(obj)
}

const useGetList = <T>(key: string, url: string, pagination?: any, filters?: any, sorter?: any) => {
  const axios = useAxios()

  const service = async () => {
    const page = { ...transformPagination(pagination) }
    const params = transformParams(Object.assign(page, filters, sorter))
    const transformRequest: AxiosRequestTransformer = (data, headers) => {}
    // console.log('--params: ', params)
    const data: T = await axios.get(`${url}`, {
      params,
      paramsSerializer: (params) => {
        return qs.stringify(params, { arrayFormat: 'repeat' })
      },
      transformRequest
    })

    return data
  }
  return useQuery(key, () => service())
}

const useGetOne = <T>(key: string, url: string, params?: any) => {
  const axios = useAxios()

  const service = async () => {
    const data: T = await axios.get(`${url}`, params)

    return data
  }
  return useQuery(key, () => service())
}

const useCreate = <T, U>(url: string) => {
  const axios = useAxios()
  return useMutation(async (params: T) => {
    const data: U = await axios.post(`${url}`, qs.stringify(params))
    return data
  })
}

const useUpdate = <T>(url: string) => {
  const axios = useAxios()
  return useMutation(async (item: T) => {
    const data: T = await axios.patch(`${url}`, item)
    return data
  })
}

const useDelete = <T>(url: string) => {
  const axios = useAxios()
  return useMutation(async (id: number) => {
    const data: T = await axios.delete(`${url}?id=${id}`)
    return data
  })
}

const useBatch = (url: string) => {
  const axios = useAxios()
  return useMutation(async (ids: number[]) => {
    const data = await axios.post(`${url}`, { idList: ids })
    return data
  })
}

export { useGetOne, useGetList, useUpdate, useCreate, useDelete, useBatch }

export default axios
