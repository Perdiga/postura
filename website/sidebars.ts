import type { SidebarsConfig } from '@docusaurus/plugin-content-docs';

const sidebars: SidebarsConfig = {
  docs: [
    'index',
    'quickstart',
    {
      type: 'category',
      label: 'Guides',
      items: ['guides/commands', 'guides/rules', 'guides/configuration', 'guides/ci'],
    },
    {
      type: 'category',
      label: 'Reference',
      items: ['reference/facts', 'reference/permissions', 'reference/reports'],
    },
    'development',
  ],
};

export default sidebars;