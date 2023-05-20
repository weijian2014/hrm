<script lang="ts" setup>
import * as echarts from "echarts/core"
import { PieChart, BarChart, LineChart } from "echarts/charts"
import {
   TitleComponent,
   TooltipComponent,
   GridComponent,
   // 数据集组件
   DatasetComponent,
   // 内置数据转换器组件 (filter, sort)
   TransformComponent,
} from "echarts/components"
import { LabelLayout, UniversalTransition } from "echarts/features"
import { CanvasRenderer } from "echarts/renderers"
import type {
   // 系列类型的定义后缀都为 SeriesOption
   PieSeriesOption,
   BarSeriesOption,
   LineSeriesOption,
} from "echarts/charts"
import type {
   // 组件类型的定义后缀都为 ComponentOption
   TitleComponentOption,
   TooltipComponentOption,
   GridComponentOption,
   DatasetComponentOption,
} from "echarts/components"
import type { ComposeOption } from "echarts/core"

// 通过 ComposeOption 来组合出一个只有必须组件和图表的 Option 类型
export type ECOption = ComposeOption<
   | PieSeriesOption
   | BarSeriesOption
   | LineSeriesOption
   | TitleComponentOption
   | TooltipComponentOption
   | GridComponentOption
   | DatasetComponentOption
>

// 注册必须的组件
echarts.use([
   PieChart,
   BarChart,
   LineChart,
   TitleComponent,
   TooltipComponent,
   GridComponent,
   DatasetComponent,
   TransformComponent,
   LabelLayout,
   UniversalTransition,
   CanvasRenderer,
])

const props = defineProps<{
   visible: boolean
   options: ECOption
}>()

const chartRef = ref()
onMounted(() => {
   console.log("CommonEcharts")
   var chart = echarts.init(chartRef.value)
   if (props.options) {
      chart.setOption(props.options)
   }

   watchEffect(() => {
      chart.setOption(props.options)
   })
})
</script>

<template>
   <!-- echarts的div一定要有高度, 否则不显示 -->
   <div v-if="visible" ref="chartRef" class="echarts-class"></div>
</template>

<style lang="scss" scoped>
.echarts-class {
   height: 200px;
}
</style>
