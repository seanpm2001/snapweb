module.exports = {
  PACKAGES: '/api/v2/packages/',
  SETTINGS: '/api/v2/packages/ubuntu-core',
  FILTERED_SNAPS: [
    'webdm',
    'ubuntu-core'
  ],
  INSTALL_STATE: {
    INSTALLED: 'installed',
    INSTALLING: 'installing',
    UNINSTALLED: 'uninstalled',
    UNINSTALLING: 'uninstalling'
  },
  INSTALL_POLL_WAIT: 16 // milliseconds
};
