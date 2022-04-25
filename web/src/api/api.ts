import { ApiList, ApiSearchParams } from '@/models/api.interface'
import { useBatch, useCreate, useGetList, useGetOne, useUpdate } from '../utils/request'

const ApiListResource = '/api/list'

export const useGetApiList = (pagination?: any, filters?: any, sorter?: any) => {
  return useGetList<API.ApiPagination>('ApiList', ApiListResource, pagination, filters, sorter)
}
