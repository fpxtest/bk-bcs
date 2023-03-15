<template>
  <bcs-navigation-menu
    item-hover-bg-color="#E1ECFF"
    item-active-bg-color="#E1ECFF"
    item-hover-color="#3A84FF"
    item-active-color="#3A84FF"
    item-hover-icon-color="#3A84FF"
    item-active-icon-color="#3A84FF"
    :unique-opened="false"
    :default-active="activeMenu.id"
    :before-nav-change="handleBeforeNavChange">
    <bcs-navigation-menu-item
      v-for="menu in activeNav.children"
      :key="menu.id"
      :id="menu.id"
      :icon="['bcs-icon', menu.icon]"
      :has-child="menu.children && !!menu.children.length"
      :disabled="disabledMenuIDs.includes(menu.id)"
      @click="handleChangeMenu(menu)">
      <span>{{ menu.title }}</span>
      <bcs-tag theme="danger" v-if="menu.tag">{{ menu.tag }}</bcs-tag>
      <template #child>
        <bcs-navigation-menu-item
          v-for="child in menu.children"
          :key="child.id"
          :id="child.id"
          :icon="['bcs-icon', child.icon]"
          :disabled="disabledMenuIDs.includes(child.id)"
          @click="handleChangeMenu(child)">
          <span>{{ child.title }}</span>
          <bcs-tag theme="danger" v-if="child.tag">{{ child.tag }}</bcs-tag>
        </bcs-navigation-menu-item>
      </template>
    </bcs-navigation-menu-item>
  </bcs-navigation-menu>
</template>
<script lang="ts">
import { defineComponent, ref, watch, computed, toRef, reactive } from '@vue/composition-api';
import $router from '@/router';
import $store from '@/store';
import { useCluster, useConfig, useProject } from '@/composables/use-app';
import menusData, { IMenu } from './menus';

export default defineComponent({
  name: 'SideMenu',
  setup() {
    const menus = ref<IMenu[]>(menusData);
    // 左侧菜单
    const activeMenu = ref<Partial<IMenu>>({});
    // 一级菜单
    const activeNav = ref<Partial<IMenu>>({});
    // 所有叶子菜单项
    const leafMenus = computed(() => flatLeafMenus(menus.value));
    // 当前路由
    const route = computed(() => toRef(reactive($router), 'currentRoute').value);

    // 共享集群禁用菜单
    const { _INTERNAL_ } = useConfig();
    const { isSharedCluster } = useCluster();
    const disabledMenuIDs = computed(() => {
      const disabledIDs: string[] = [];
      if (isSharedCluster.value) {
        disabledIDs.push(...[
          'DAEMONSET',
          'STORAGE',
          'RBAC',
          'HPA',
          'CRD',
          'CUSTOMOBJECT',
        ]);
      }
      if (_INTERNAL_.value) {
        disabledIDs.push('CLOUDTOKEN');
      }
      return disabledIDs;
    });

    // 设置当前菜单ID
    watch(route, () => {
      // 路由上配置了菜单ID或者路由名称与当前子菜单项路由名称一致
      const menu = leafMenus.value
        .find(item => item.route === route.value.name || item.id === route.value.meta?.menuId);

      if (!menu) {
        console.warn(`current route ${route.value.name} has no matched menuId`);
      } else {
        activeMenu.value = menu || {};
        activeNav.value = menu?.root || {};
        $store.commit('updateCurSideMenu', activeMenu.value);
      }
    }, { immediate: true });

    // 扁平化子菜单
    function flatLeafMenus(menus: IMenu[], root?: IMenu) {
      const data: IMenu[] = [];
      for (const item of menus) {
        const rootMenu = root ?? item;
        if (item.children?.length) {
          data.push(...flatLeafMenus(item.children, rootMenu));
        } else {
          data.push({
            root: rootMenu,
            ...item,
          });
        }
      }
      return data;
    };
    // 切换菜单
    const { projectCode } = useProject();
    const handleBeforeNavChange = () => false;
    const handleChangeMenu = (item: IMenu) => {
      if (route.value.name === item.route) return;

      if (item.id === 'MONITOR') {
        window.open(`${window.BKMONITOR_HOST}/?space_uid=bkci__${projectCode.value}#/k8s`);
      } else {
        $router.push({
          name: item.route || item.children?.[0]?.route || '404',
          params: {
            projectCode: $store.getters.curProjectCode,
          },
        });
      }
    };

    return {
      activeMenu,
      activeNav,
      disabledMenuIDs,
      handleChangeMenu,
      handleBeforeNavChange,
    };
  },
});
</script>
<style lang="postcss" scoped>
>>> .navigation-sbmenu.is-disabled {
  cursor: not-allowed;
  opacity: .3;
}
</style>