<template>
  <div :class="{'has-logo':true}">
    <logo :collapse="isCollapse" />
    <el-scrollbar wrap-class="scrollbar-wrapper" :style="{'background':this.GLOBAL.THEME_COLOR?this.GLOBAL.THEME_COLOR:''}">
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :background-color="this.GLOBAL.THEME_COLOR?this.GLOBAL.THEME_COLOR:variables.menuBg"
        :text-color="this.GLOBAL.THEME_COLOR?'#dfb5b6':variables.menuText"
        :unique-opened="false"
        :active-text-color="this.GLOBAL.THEME_COLOR?'#fff':variables.menuActiveText"
        :collapse-transition="false"
        mode="vertical"
      >
        <div v-for="route in routes" :key="route.path">
          <sidebar-item
            v-if="route.path === '/cloudInterconnection' && isPermission === 'no' ? false : true"
            :key="route.path"
            :item="route"
            :base-path="route.path"
          />
        </div>
        <!-- <sidebar-item v-for="route in routes" :key="route.path" :item="route" :base-path="route.path" /> -->
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Logo from './Logo'
import SidebarItem from './SidebarItem'
import variables from '@/styles/variables.scss'
import { getUserConfig } from '@/api/Home'

export default {
  components: { SidebarItem, Logo },
  data(){
    return {
      isPermission: 'no'
    }
  },
  created() {
    this.getUserConfig()
  },
  methods: {
    getUserConfig() {
      getUserConfig().then(response => {
        if (response.success) {
          this.isPermission = response.data.config&&response.data.config.jointCloudPermission?response.data.config.jointCloudPermission:'no'
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    }
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    routes() {
      return this.$router.options.routes
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      // if set path, the sidebar will highlight the path you set
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  }
}
</script>
