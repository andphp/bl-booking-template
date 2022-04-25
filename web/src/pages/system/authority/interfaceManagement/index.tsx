
import React, { useEffect, useRef, useState } from 'react'
import { PlusOutlined, EllipsisOutlined, SyncOutlined } from '@ant-design/icons'
import { Button, Tag, Space, Menu, Dropdown, PaginationProps } from 'antd'
import type { ProColumns, ActionType } from '@ant-design/pro-table'
import ProTable, { TableDropdown } from '@ant-design/pro-table'
import { useGetApiList } from '@/api/api'
import { ApiItem, ApiList, ApiSearchParams } from '@/models/api.interface'
import { superFilter } from '@/utils/helper'
import { useUpdateEffect } from 'ahooks'

const columns: ProColumns<ApiItem>[] = [
  {
    dataIndex: 'id',
    valueType: 'indexBorder',
    width: 48
  },
  {
    title: '名称',
    dataIndex: 'name',
    copyable: true,
    ellipsis: true,
    tip: '标题过长会自动收缩',
    formItemProps: {
      rules: [
        {
          required: true,
          message: '此项为必填项'
        }
      ]
    }
  },
  {
    title: '路由路径',
    dataIndex: 'path',
    copyable: true,
    ellipsis: true,
    tip: '标题过长会自动收缩',
    formItemProps: {
      rules: [
        {
          required: true,
          message: '此项为必填项'
        }
      ]
    }
  },
  {
    title: '描述',
    dataIndex: 'description',
    ellipsis: true,
    tip: '标题过长会自动收缩'
  },
  {
    title: '路由组',
    dataIndex: 'apiGroup',
    ellipsis: true,
    tip: '标题过长会自动收缩'
  },
  {
    title: '请求方法',
    dataIndex: 'method',
    copyable: true,
    ellipsis: true,
    tip: '标题过长会自动收缩'
  },
  {
    title: '创建时间',
    key: 'showTime',
    dataIndex: 'createdAt',
    valueType: 'dateTime',
    sorter: true,
    hideInSearch: true
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    valueType: 'dateRange',
    hideInTable: true,
    search: {
      transform: (value) => {
        return {
          startTime: value[0],
          endTime: value[1]
        }
      }
    }
  },
  {
    title: '操作',
    valueType: 'option',
    render: (text, record, _, action) => [
      <a
        key='editable'
        onClick={() => {
          action?.startEditable?.(record.id)
        }}
      >
        编辑
      </a>
    ]
  }
]

const InterfaceManagement: React.FC = () => {
  const actionRef = useRef<ActionType>()

  const [ApiList, setApiList] = useState<ApiList>()
  const [filters, setFilters] = useState<ApiList>()
  const [pagination, setPagination] = useState<Partial<PaginationProps>>({
    current: 1,
    pageSize: 10,
    total: 0
  })
  const { data: apiListData, refetch: apiRefetch, isFetching: apiFetching, isLoading: apiLoading } = useGetApiList(pagination, filters)

  useEffect(() => {
    setApiList(apiListData?.data)
    setPagination({
      ...pagination,
      total: apiListData?.total,
      showQuickJumper: true
    })
  }, [apiListData])

  useUpdateEffect(() => {
    apiRefetch()
  }, [pagination.current, pagination.pageSize, filters])

  return (
    <ProTable<ApiItem>
      columns={columns}
      actionRef={actionRef}
      rowKey='id'
      options={{ reload: false }}
      editable={{
        type: 'multiple'
      }}
      dataSource={ApiList}
      pagination={pagination}
      onChange={(pagination, filters, sorter) => {
        setPagination(pagination)
      }}
      search={{
        defaultCollapsed: false,
        optionRender: ({ searchText, resetText }, { form }) => [
          <Button
            key='search'
            type='primary'
            onClick={() => {
              // form?.submit();
              // console.log('search submit', form?.getFieldsValue())
              setFilters(form?.getFieldsValue())
            }}
          >
            {searchText}
          </Button>,
          <Button
            key='reset'
            onClick={() => {
              form?.resetFields()
            }}
          >
            {resetText}
          </Button>
        ]
      }}
      loading={apiFetching || apiLoading}
      form={{
        // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
        syncToUrl: (values, type) => {
          if (type === 'get') {
            return {
              ...values,
              created_at: [values.startTime, values.endTime]
            }
          }
          return values
        }
      }}
      dateFormatter='string'

      headerTitle={
        <Button key='button' icon={<SyncOutlined />}>
        同步/更新
        </Button>
      }
    />
  )
}

export default InterfaceManagement
