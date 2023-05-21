<script lang="ts" setup>
import EmployeeAddOrEditVue from "@/components/employee/EmployeeAddOrEdit.vue"
import EchartsVue, { ECOption } from "@/components/echarts/CommonEcharts.vue"
import { employeeListApi, employeeSearchApi, employeeDeleteApi } from "@/utils/employee"
import { postListApi } from "@/utils/post"
import type { CheckboxValueType } from "element-plus"
import { useSettings, useData, convertForExport } from "./index"
import moment from "moment"
import { getAgeByBirthday } from "@/utils/common"
import XLSXS from "xlsx-js-style"

const { table, columns, pagination, checkList } = useSettings()

const {
   tableRef,
   tableData,
   tableRows,
   isButtonDisabled,
   searchInputValue,
   isAddOrEditShow,
   addOrEditTitle,
   addOrEditData,
   selections,
   isLoading,
   posts,
} = useData()

const refresh = async () => {
   isLoading.value = true

   await postListApi()
      .then((res) => {
         if (res.code === 200) {
            console.log("postListApi", res)
            posts.value = res.data.posts
         }
      })
      .catch((res) => {
         console.log(res)
         return new Promise(() => {})
      })

   await employeeListApi()
      .then((res) => {
         if (res.code === 200) {
            console.log("employeeListApi", res)
            tableData.value = res.data.employees
            tableRows.value = res.data.total
            currentPage.value = 1
         }
      })
      .catch((res) => {
         console.log(res)
         return new Promise(() => {})
      })
      .finally(() => {
         isLoading.value = false
      })
}

refresh()

const deleteAndRefresh = async (ids: number[]) => {
   for (let i = 0; i < ids.length; ++i) {
      await employeeDeleteApi(ids[i])
         .then((res) => {
            if (res.code === 200) {
               console.log(res)
            }
         })
         .catch((res) => {
            console.log(res)
            isButtonDisabled.value = false
            return new Promise(() => {})
         })
   }

   await refresh()
}

// 可显示的列
const visibleColumns = computed(() => {
   return columns.filter((column: { visible: boolean }) => column.visible)
})

// 查询输入变化时
watch(
   () => searchInputValue.value,
   async (newValue) => {
      if (newValue && newValue.length != 0) {
         await employeeSearchApi({ key: newValue })
            .then((res) => {
               if (res.code === 200) {
                  tableData.value = res.data.employees
                  tableRows.value = res.data.total
                  currentPage.value = 1
               }
            })
            .catch((err) => {
               console.log(err)
               return new Promise(() => {})
            })
      } else {
         await refresh()
      }
   }
)

////// 表头工具栏
let lastCheckList = checkList.value as CheckboxValueType[]
// 列筛选
const handleCheckedChange = (values: CheckboxValueType[]) => {
   // values为所有被选中的列
   // 求出values与lastCheckedLables的差集--取消选中的列
   const diffA = values.concat(lastCheckList).filter((v) => !values.includes(v))

   // 求出lastCheckedLables与values的差集--新增选中的列
   var diffB = lastCheckList.concat(values).filter((v) => !lastCheckList.includes(v))

   diffA.forEach((item1) => {
      columns.forEach((item2) => {
         if (item2.label === item1) {
            item2.visible = false
         }
      })
   })

   diffB.forEach((item1) => {
      columns.forEach((item2) => {
         if (item2.label === item1) {
            item2.visible = true
         }
      })
   })

   lastCheckList = values
}

// 表格缩放
const tableSize = ref("正常")
watch(
   () => tableSize.value,
   (newValue) => {
      if (newValue == "缩小") {
         table.size = "small"
      } else if (newValue === "正常") {
         table.size = "default"
      } else {
         table.size = "large"
      }
   }
)

// 统计图表
const isChartVisable = ref(true)
const echartsVisableValue = ref("显示图表")
watch(
   () => echartsVisableValue.value,
   (newValue) => {
      if (newValue === "隐藏图表") {
         isChartVisable.value = false
      } else {
         // 显示图表
         isChartVisable.value = true
      }
   }
)

////// 表格

// 勾选表格的行时
const handleSelectChange = (rows: Employee[]) => {
   // rows为所有选中的行
   console.log("handleSelectChange", rows)
   selections.value = rows
   isButtonDisabled.value = rows.length === 0
}

const handleRowClick = (row: Employee) => {
   // console.log("handleRowClick", row)
}

// 新增或导入
const handleAddOrImport = () => {
   console.log("handleAddOrImport")
   addOrEditTitle.value = "新增或导入员工"
   addOrEditData.value = {} as Employee
   isAddOrEditShow.value = true
}

function setExcelStyle(
   sheet: XLSXS.WorkSheet,
   excelHeaders: any[],
   excelBodys: any[],
   excelColumnsWidth: any[],
   excelColumnsAlignment: any[]
) {
   // 设置列宽度
   let colsWidth: XLSXS.ColInfo[] = []
   excelColumnsWidth.forEach((w) => {
      const col = { wpx: w }
      colsWidth.push(col)
   })
   sheet["!cols"] = colsWidth

   // 设置行高度
   const rows = [{ hpx: 32 }, { hpx: 20 }]
   sheet["!rows"] = rows

   const borderAll = {
      //单元格外侧框线
      color: { auto: 1 },
      top: {
         style: "thin",
      },
      bottom: {
         style: "thin",
      },
      left: {
         style: "thin",
      },
      right: {
         style: "thin",
      },
   }

   const style = {
      // 表头样式
      hs: {
         border: borderAll,
         font: { sz: 11, color: { rgb: "409EFF" }, bold: true },
         alignment: { horizontal: "center", vertical: "center", wrapText: true },
         // 背景色
         // fill: { bgColor: { indexed: 64 }, fgColor: { rgb: "FFFF00" } },
      },
      // 内容样式
      bs: {
         border: borderAll,
         font: { sz: 10 },
         alignment: { horizontal: "center", vertical: "center", wrapText: true },
      },
      // 注释行样式
      ts: {
         font: { sz: 10, bold: true },
         alignment: { vertical: "center", wrapText: true },
         fill: { bgColor: { indexed: 64 }, fgColor: { rgb: "00B050" } },
      },
   }

   // 设置边框, 对齐, 字体等样式
   for (const key in sheet) {
      // 第一行的单元格, 即表头
      if (Number(key.slice(1)) == 1 && sheet[key].t) {
         sheet[key].s = style.hs
      } else if (sheet[key].t) {
         // 数据行单元格
         sheet[key].s = style.bs
      }
   }
}

// 批量导出
const handleExport = () => {
   if (selections.value.length === 0) {
      return
   }

   let names: string = ""
   selections.value.forEach((item, i) => {
      names += item.name
      if (i != selections.value.length - 1) {
         names += ", "
      }
   })
   ElMessageBox.confirm(
      h("p", null, [
         h("span", null, "确认导出姓名 "),
         h("i", { style: "color: red" }, names),
         h("span", null, " 共" + selections.value.length + "条记录?"),
      ]),
      "员工导出确认",
      {
         confirmButtonText: "确认",
         cancelButtonText: "取消",
         type: "info",
      }
   )
      .then(() => {
         // excelData包含表头和行数据
         const { excelHeaders, excelBodys, excelColumnsWidth, excelColumnsAlignment } = convertForExport(
            selections.value
         )
         console.log("handleExport", excelHeaders, excelBodys, excelColumnsWidth, excelColumnsAlignment)

         // 创建工作簿
         const book = XLSXS.utils.book_new()

         // 创建工作表, skipHeader=true, 因为excelData中已经包含表头
         const sheet = XLSXS.utils.json_to_sheet([excelHeaders, ...excelBodys], { skipHeader: true })

         // 设置表格样式
         setExcelStyle(sheet, excelHeaders, excelBodys, excelColumnsWidth, excelColumnsAlignment)

         // 将工作表放入工作簿中
         XLSXS.utils.book_append_sheet(book, sheet, "员工信息")

         // 生成文件并下载
         const excelFileName = "员工信息_" + moment(new Date()).format("YYYY-MM-DD") + ".xlsx"
         XLSXS.writeFile(book, excelFileName)

         // 清除行选中
         tableRef.value.clearSelection()
      })
      .catch(() => {
         ElMessage.info("操作已取消")
         isButtonDisabled.value = false
         return new Promise(() => {})
      })
}

// 批量删除
const handleBatchDelete = () => {
   console.log("handleBatchDelete", selections)
   if (selections.value.length === 0) {
      isButtonDisabled.value = true
      return
   }

   isButtonDisabled.value = true
   let names: string = ""
   let ids: number[] = []
   selections.value.forEach((item, i) => {
      names += item.name
      ids.push(item.id)
      if (i != selections.value.length - 1) {
         names += ", "
      }
   })

   ElMessageBox.confirm(
      h("p", null, [
         h("span", null, "确认删除姓名 "),
         h("i", { style: "color: red" }, names),
         h("span", null, " 共" + selections.value.length + "条记录?"),
      ]),
      "员工删除确认",
      {
         confirmButtonText: "确认",
         cancelButtonText: "取消",
         type: "warning",
      }
   )
      .then(() => {
         deleteAndRefresh(ids)

         selections.value = [] as Employee[]
         isButtonDisabled.value = true
         ElMessage.success("删除成功")
      })
      .catch(() => {
         ElMessage.info("操作已取消")
         isButtonDisabled.value = false
         return new Promise(() => {})
      })
}

// 行编辑
const handleEdit = (index: number, row: Employee | undefined) => {
   console.log("handleEdit", index, row)
   if (row) {
      addOrEditData.value = row
      console.log("handleEdit", addOrEditData.value)
      addOrEditTitle.value = "修改员工"
      isAddOrEditShow.value = true
   }
}

// 行删除
const handleDelete = async (index: number, row: Employee) => {
   console.log("handleDelete", index, row)
   await ElMessageBox.confirm(
      h("p", null, [
         h("span", null, "确认删除姓名 "),
         h("i", { style: "color: red" }, row.name),
         h("span", null, " 的记录?"),
      ]),
      "员工删除确认",
      {
         confirmButtonText: "确认",
         cancelButtonText: "取消",
         type: "warning",
      }
   ).catch(() => {
      ElMessage.info("操作已取消")
      return new Promise(() => {})
   })

   const ids = [row.id]
   deleteAndRefresh(ids)
   ElMessage.success("删除成功")
}

// 刷新
const handleRefresh = async () => {
   console.log("handleRefresh")
   await refresh()
}

////// Add组件

// AddVue组件发送的保存事件
const handleSave = async (message: string) => {
   isAddOrEditShow.value = false
   console.log("handleSave", message, addOrEditData.value)

   // 从数据库中读取最新的数据
   await refresh()
   ElMessage.success(message)
}

// AddVue组件发送的消息事件
const handleCancel = (message: string) => {
   isAddOrEditShow.value = false
   console.log("handleCancel", message, addOrEditData.value)
   ElMessage.info(message)
}

// 分页
const currentPage = ref(1)
const pageSize = ref(10)
const handleCurrentChange = (value: number) => {
   console.log("handleCurrentChange", value)
   currentPage.value = value
}
const handleSizeChange = (value: number) => {
   console.log("handleSizeChange", value)
   pageSize.value = value
   currentPage.value = 1
}

watch(
   () => tableData.value,
   (newValue) => {
      // 性别
      genderOptionData.value[0].value = 0
      genderOptionData.value[1].value = 0

      // 年龄
      ageOptionData.value.forEach((v, i) => {
         ageOptionData.value[i].value = 0
      })

      // 岗位
      if (postOptionData.value.x.length === 0) {
         posts.value.forEach((p) => {
            postOptionData.value.x.push(p.name)
            postOptionData.value.y.push(0)
         })
      } else {
         postOptionData.value.y.map((v, i) => {
            postOptionData.value.y[i] = 0
         })
      }

      // 工资
      salaryOptionData.value.forEach((v, i) => {
         salaryOptionData.value[i].value = 0
      })

      var index = 0
      if (tableData.value) {
         tableData.value.forEach((e) => {
            // 性别
            if (e.gender === "男") {
               genderOptionData.value[0].value += 1
            } else {
               genderOptionData.value[1].value += 1
            }

            // 年龄
            const dateStr = moment(e.birthday).format("YYYY-MM-DD")
            const age = getAgeByBirthday(dateStr)
            if (age < 25) {
               index = 0
            } else if (age >= 25 && age < 30) {
               index = 1
            } else if (age >= 30 && age < 35) {
               index = 2
            } else if (age >= 35 && age < 40) {
               index = 3
            } else if (age >= 40 && age < 45) {
               index = 4
            } else {
               index = 5
            }
            ageOptionData.value[index].value += 1

            // 岗位
            index = postOptionData.value.x.indexOf(e.post)
            if (index !== -1) {
               postOptionData.value.y[index] += 1
            }

            // 工资
            if (e.salary < 4000) {
               index = 0
            } else if (e.salary >= 4000 && e.salary < 5000) {
               index = 1
            } else if (e.salary >= 5000 && e.salary < 6000) {
               index = 2
            } else if (e.salary >= 6000 && e.salary < 7000) {
               index = 3
            } else if (e.salary >= 7000 && e.salary < 8000) {
               index = 4
            } else {
               index = 5
            }
            salaryOptionData.value[index].value += 1
         })
      }

      console.log("genderOptionData", genderOptionData.value)
      console.log("ageOptionData", ageOptionData.value)
      console.log("postOptionData", postOptionData.value)
      console.log("salaryOptionData", salaryOptionData.value)
   }
)

// 性别
const genderOptionData = ref([
   { name: "男", value: 0 },
   { name: "女", value: 0 },
])
const genderOption: ECOption = {
   title: {
      text: "性别结构",
      left: "center",
   },
   tooltip: {
      trigger: "item",
      formatter: "{b}: {c}人, 占{d}%",
   },
   series: {
      type: "pie",
      radius: [30, 60],
      itemStyle: {
         borderColor: "#fff",
         borderWidth: 1,
      },
      label: {
         alignTo: "labelLine",
         formatter: "{name|{b}}\n{time|{c} 人}",
         minMargin: 5,
         edgeDistance: 10,
         lineHeight: 15,
         rich: {
            time: {
               fontSize: 10,
               color: "#999",
            },
         },
      },
      labelLine: {
         length: 15,
         length2: 0,
         maxSurfaceAngle: 80,
      },
      data: genderOptionData.value,
   },
}

// 年龄
const ageOptionData = ref<{ value: number; name: string }[]>([
   { value: 0, name: "25岁以下" },
   { value: 0, name: "25~30岁" },
   { value: 0, name: "30~35岁" },
   { value: 0, name: "35~40岁" },
   { value: 0, name: "40~45岁" },
   { value: 0, name: "45岁以上" },
])
const ageOption: ECOption = {
   title: {
      text: "年龄结构",
      left: "center",
   },
   tooltip: {
      trigger: "item",
      formatter: "{b}: {c}人, 占{d}%",
   },
   legend: {
      show: false,
      orient: "vertical",
      left: "left",
   },
   series: [
      {
         name: "Access From",
         type: "pie",
         radius: "60%",
         data: ageOptionData.value,
         label: {
            alignTo: "labelLine",
            formatter: "{name|{b}}\n{time|{c} 人}",
            minMargin: 5,
            edgeDistance: 10,
            lineHeight: 15,
            rich: {
               time: {
                  fontSize: 10,
                  color: "#999",
               },
            },
         },
         emphasis: {
            itemStyle: {
               shadowBlur: 20,
               shadowOffsetX: 0,
               shadowColor: "rgba(0, 0, 0, 0.5)",
            },
         },
      },
   ],
}

// 岗位
const postOptionData = ref<{ x: string[]; y: number[] }>({ x: [], y: [] })
const postOption: ECOption = {
   title: {
      text: "岗位结构",
      left: "center",
   },
   tooltip: {
      trigger: "item",
      formatter: "{b}: {c}人, 占{d}%",
   },
   xAxis: {
      type: "category",
      data: postOptionData.value.x,
   },
   yAxis: {
      type: "value",
   },
   series: [
      {
         data: postOptionData.value.y,
         type: "bar",
         label: {
            show: true,
            formatter: "{c}人", // 显示数值
            position: "top", // 数值在柱体上方
         },
         itemStyle: {
            // 定制显示（按顺序）
            color: function (params) {
               var colorList = ["#5c7bd9", "#91cc75", "#ffdc60", "#ee6666", "#73c0de", "#3ba272"]
               return colorList[params.dataIndex]
            },
         },
         labelLine: {
            length: 15,
            length2: 0,
         },
      },
   ],
}

// 工资
const salaryOptionData = ref<{ value: number; name: string }[]>([
   { value: 0, name: "4000元以下" },
   { value: 0, name: "4000~5000元" },
   { value: 0, name: "5000~6000元" },
   { value: 0, name: "6000~7000元" },
   { value: 0, name: "7000~8000元" },
   { value: 0, name: "8000元以上" },
])
const salaryOption: ECOption = {
   title: {
      text: "工资结构",
      left: "center",
   },
   tooltip: {
      trigger: "item",
      formatter: "{b}: {c}人, 占{d}%",
   },
   series: [
      {
         name: "Nightingale Chart",
         type: "pie",
         radius: [20, 60],
         center: ["50%", "50%"],
         roseType: "area",
         itemStyle: {
            borderRadius: 8,
         },
         label: {
            alignTo: "labelLine",
            formatter: "{name|{b}}\n{time|{c} 人}",
            minMargin: 5,
            edgeDistance: 10,
            lineHeight: 15,
            rich: {
               time: {
                  fontSize: 10,
                  color: "#999",
               },
            },
         },
         data: salaryOptionData.value,
      },
   ],
}
</script>

<template>
   <div>
      <!-- 表头工具 -->
      <el-row class="mb-2">
         <el-col :span="5">
            <el-button type="primary" @click="handleAddOrImport">
               <IEpPlus />
               <span>新增或导入</span>
            </el-button>
            <el-button type="primary" @click="handleExport" :disabled="isButtonDisabled">
               <IEpDownload />
               <span>批量导出</span>
            </el-button>
            <el-button type="primary" @click="handleBatchDelete" :disabled="isButtonDisabled">
               <IEpDelete />
               <span>批量删除</span>
            </el-button>
         </el-col>
         <el-col :span="6">
            <!-- v-model.trim 去掉左右两侧的空格 -->
            <el-input v-model.trim="searchInputValue" placeholder="" clearable>
               <template #suffix> </template>
               <template #prepend
                  >模糊搜索<el-icon><IEpSearch /></el-icon
               ></template>
            </el-input>
         </el-col>
         <el-col class="ml-2" :span="1">
            <el-button type="primary">
               <IEpRefresh />
               <span style="vertical-align: middle" @click="handleRefresh">刷新</span>
            </el-button>
         </el-col>
         <el-col class="ml-8" :span="1">
            <el-popover placement="right-end" title="列筛选" :width="300" trigger="click">
               <template #reference>
                  <el-button type="primary" style="vertical-align: middle"><IEpFinished />列筛选</el-button>
               </template>
               <el-checkbox-group v-model="checkList" @change="handleCheckedChange" :min="6" :max="17">
                  <el-checkbox
                     v-for="item in columns"
                     :label="item.label"
                     :checked="item.visible"
                     :disabled="item.disable"
                     v-show="item.prop != 'id'"
                     >{{ item.label }}</el-checkbox
                  >
               </el-checkbox-group>
            </el-popover>
         </el-col>
         <el-col class="ml-12" :span="3">
            <el-radio-group v-model="tableSize">
               <el-radio-button label="缩小" />
               <el-radio-button label="正常" />
               <el-radio-button label="放大" />
            </el-radio-group>
         </el-col>
         <el-col :span="3">
            <el-radio-group v-model="echartsVisableValue">
               <el-radio-button label="隐藏图表" />
               <el-radio-button label="显示图表" />
            </el-radio-group>
         </el-col>
      </el-row>
      <!-- 表格, table-layout为auto时表头无法固定 -->
      <el-table
         ref="tableRef"
         style="width: 100%"
         :data="tableData.slice((currentPage - 1) * pageSize, currentPage * pageSize)"
         :border="table.border"
         :fit="table.fit"
         :size="table.size"
         :table-layout="table.table_layout"
         :empty-text="table.empty_text"
         :highlight-current-row="table.highlight_current_row"
         :row-key="table.row_key"
         :default-sort="{ prop: 'id', order: 'descending' }"
         v-loading="isLoading"
         @selection-change="handleSelectChange"
         @row-click="handleRowClick">
         <el-table-column type="selection" align="center" />
         <el-table-column
            v-for="(column, index) in visibleColumns"
            :key="column.prop"
            :prop="column.prop"
            :label="column.label"
            :sortable="column.sortable"
            :fixed="column.fixed"
            :align="column.align"
            :index="index">
            <!-- v-html和formatter不能同时使用, 另外formatter不支持html样式 -->
            <template #default="scope">
               <span v-html="column.formatter(scope.row, column.prop, searchInputValue)" />
            </template>
         </el-table-column>
         <el-table-column fixed="right" width="145" label="操作" align="center">
            <template #default="scope">
               <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
               <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
         </el-table-column>
      </el-table>
      <el-pagination
         class="mt-2"
         v-model:current-page="currentPage"
         v-model:page-size="pageSize"
         :background="pagination.background"
         :hide-on-single-page="pagination.hide_on_single_page"
         :page-sizes="pagination.page_sizes"
         :layout="pagination.layout"
         :total="tableRows"
         :prev-text="pagination.prev_text"
         :next-text="pagination.next_text"
         @size-change="handleSizeChange"
         @current-change="handleCurrentChange" />
   </div>
   <el-row v-if="isChartVisable" class="mt-6">
      <el-col :span="6">
         <EchartsVue :options="genderOption"></EchartsVue>
      </el-col>
      <el-col :span="6">
         <EchartsVue :options="ageOption"></EchartsVue>
      </el-col>
      <el-col :span="6">
         <EchartsVue :options="postOption"></EchartsVue>
      </el-col>
      <el-col :span="6">
         <EchartsVue :options="salaryOption"></EchartsVue>
      </el-col>
   </el-row>

   <EmployeeAddOrEditVue
      :isShow="isAddOrEditShow"
      :title="addOrEditTitle"
      :formData="addOrEditData"
      :postData="posts"
      @save="handleSave"
      @cancel="handleCancel">
   </EmployeeAddOrEditVue>
</template>

<style lang="scss" scoped></style>
