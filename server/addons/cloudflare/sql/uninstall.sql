-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和在SQL文件：/Volumes/NetAC/www/hotgo/server/addons/cloudflare/sql/uninstall.sql
-- Version: 2.9.3
-- Date: 2023-10-14 17:57:19
-- Link https://github.com/bufanyun/hotgo

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `hotgo`
--

-- --------------------------------------------------------
-- 菜单目录
delete from `hg_admin_menu` where `name` = 'account' and `path` = '/addons/account';
-- 菜单页面
-- 列表
delete from `hg_admin_menu` where `permissions` = '/addons/cloudflare/account/list';

-- 详情
delete from `hg_admin_menu` where `permissions` = '/addons/cloudflare/account/view';

-- 菜单按钮

-- 编辑
delete from `hg_admin_menu` where `permissions` = '/addons/cloudflare/account/edit';



-- 删除
delete from `hg_admin_menu` where `permissions` = '/addons/cloudflare/account/delete';




-- 导出
delete from `hg_admin_menu` where `permissions` = '/addons/cloudflare/account/export';

COMMIT;