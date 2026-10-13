import type { Config } from '@docusaurus/types';

const config: Config = {
  title: 'postura',
  tagline: 'A security baseline you own.',
  favicon: 'img/favicon.svg',
  url: 'https://jackchuka.github.io',
  baseUrl: '/postura/',
  organizationName: 'jackchuka',
  projectName: 'postura',
  onBrokenLinks: 'throw',
  markdown: {
    hooks: {
      onBrokenMarkdownLinks: 'throw',
    },
  },
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  presets: [
    [
      'classic',
      {
        docs: {
          routeBasePath: '/',
          sidebarPath: './sidebars.ts',
          editUrl: 'https://github.com/jackchuka/postura/tree/main/website/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      },
    ],
  ],
  themeConfig: {
    navbar: {
      title: 'postura',
      items: [
        { type: 'docSidebar', sidebarId: 'docs', position: 'left', label: 'Documentation' },
        {
          href: 'https://github.com/jackchuka/postura',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'light',
      links: [
        {
          title: 'Learn',
          items: [
            { label: 'Quickstart', to: '/quickstart' },
            { label: 'Rules and CEL', to: '/guides/rules' },
            { label: 'Configuration', to: '/guides/configuration' },
          ],
        },
        {
          title: 'Reference',
          items: [
            { label: 'Fact fields', to: '/reference/facts' },
            { label: 'Permissions', to: '/reference/permissions' },
            { label: 'GitHub repository', href: 'https://github.com/jackchuka/postura' },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} postura contributors. MIT License.`,
    },
    prism: {
      additionalLanguages: ['bash', 'json', 'yaml'],
    },
  },
};

export default config;