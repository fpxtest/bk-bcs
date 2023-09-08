<script lang="ts" setup>
  import { ref, computed } from 'vue'
  import { ITemplateBoundByAppData } from '../../../../../../../../../types/config';
  import { IAllPkgsGroupBySpaceInBiz, IPkgTreeItem } from '../../../../../../../../../types/template';
  import SearchInput from '../../../../../../../../components/search-input.vue';

  interface ISpaceTreeItem extends IPkgTreeItem {
    children: IPkgTreeItem[];
  }

  const props = defineProps<{
    bkBizId: string;
    pkgList: IAllPkgsGroupBySpaceInBiz[];
    imported: ITemplateBoundByAppData[]; // 编辑状态下已经选中的套餐，该列表下数据不可取消
    value: ITemplateBoundByAppData[]; // 当前选中的套餐信息
  }>()

  const emits = defineEmits(['change'])

  const searchStr = ref('')

  const pkgTreeData = computed(() => {
    let list: IAllPkgsGroupBySpaceInBiz[] = []
    if (searchStr.value) {
      props.pkgList.forEach(item => {
        const pkgs = item.template_sets.filter(pkg => pkg.template_set_name.toLocaleLowerCase().includes(searchStr.value.toLocaleLowerCase()))
        if (pkgs.length > 0) {
          list.push({ ...item, template_sets: pkgs })
        }
      })
    } else {
      list = props.pkgList
    }
    return list.map(item => {
      const { template_space_id, template_space_name, template_sets } = item
      const nodeId = `space_${template_space_id}`
      let checked = false
      let indeterminate = false

      if (template_sets.length > 0) {
        checked = template_sets.every(pkgItem => isPkgNodeChecked(pkgItem.template_set_id))
        if (!checked) {
          indeterminate = template_sets.some(pkgItem => isPkgNodeChecked(pkgItem.template_set_id))
        }
      }

      const group: ISpaceTreeItem = {
        id: template_space_id,
        nodeId,
        name: template_space_name === 'default_space' ? '默认空间' : template_space_name,
        children: [],
        checked,
        indeterminate,
        disabled: false,
      }

      template_sets.forEach(pkg => {
        group.children.push({
          id: pkg.template_set_id,
          nodeId: `pkg_${pkg.template_set_id}`,
          name: pkg.template_set_name,
          checked: isPkgNodeChecked(pkg.template_set_id),
          disabled: isPkgImported(pkg.template_set_id) || pkg.template_ids.length === 0,
          indeterminate: false
        })
      })
      return group
    })
  })

  const handleSearch = () => {}

  const handleNodeCheckChange = (node: ISpaceTreeItem, val: boolean) => {
    const list = props.value.slice()
    if (node.children) { // 空间节点
      const pkgNodes = node.children
      if (val) {
        pkgNodes.forEach(pkg => {
          if (!pkg.disabled && !isPkgNodeChecked(pkg.id)) {
            list.push({
              template_set_id: pkg.id,
              template_revisions: []
            })
          }
        })
      } else {
        pkgNodes.forEach(pkg => {
          if (isPkgImported(pkg.id)) return
          const index = list.findIndex(item => item.template_set_id === pkg.id)
          if (index > -1) {
            list.splice(index, 1)
          }
        })
      }
    } else { // 套餐节点
      if (isPkgImported(node.id)) return
      if (val) {
        if (!isPkgNodeChecked(node.id)) {
          list.push({
            template_set_id: node.id,
            template_revisions: []
          })
        }
      } else {
        const index = list.findIndex(item => item.template_set_id === node.id)
        if (index > -1) {
          list.splice(index, 1)
        }
      }
    }
    console.log(list)
    emits('change', list)
  }

  const isPkgImported = (id: number) => {
    return props.imported.some(pkg => {
      return pkg.template_set_id === id
    })
  }

  const isPkgNodeChecked = (id: number) => {
    return isPkgImported(id) || props.value.some(pkg => pkg.template_set_id === id)
  }

</script>
<template>
  <h4 class="title">从配置模板导入</h4>
  <SearchInput v-model="searchStr" placeholder="搜索模板套餐" @search="handleSearch" />
  <div class="packages-tree">
    <bk-tree
      ref="treeRef"
      label="name"
      node-key="nodeId"
      :empty-text="searchStr === '' ? '暂无搜索结果' : '暂无数据'"
      :expand-all="searchStr !== ''"
      :data="pkgTreeData"
      :show-node-type-icon="false">
      <template #node="node">
        <div class="node-item-wrapper">
          <bk-checkbox
            size="small"
            :model-value="node.checked"
            :disabled="node.disabled"
            :indeterminate="node.indeterminate"
            @change="handleNodeCheckChange(node, $event)" />
          <div class="node-name-text">{{ node.name }}</div>
          <span v-if="node.children" class="num">({{ node.children.length }})</span>
        </div>
      </template>
    </bk-tree>
  </div>
</template>
<style lang="scss" scoped>
  .title {
    margin: 0 0 25px;
    padding: 0 24px;
    line-height: 22px;
    font-size: 14px;
    font-weight: 400;
    color: #313238;
  }
  .search-input {
    padding: 0 24px;
  }
  .packages-tree {
    margin-top: 8px;
    padding: 0 24px;
    height: calc(100% - 87px);
    overflow: auto;
    .node-item-wrapper {
      display: flex;
      align-items: center;
    }
  }
  .node-item-wrapper {
    display: flex;
    align-items: center;
    .node-name-text {
      padding: 0 4px;
      font-size: 12px;
      color: #63656e;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
    }
    .num {
      flex-shrink: 0;
      font-size: 12px;
      color: #63656e;
    }
  }
</style>