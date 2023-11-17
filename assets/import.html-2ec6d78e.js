import{_ as a,o as n,c as s,e}from"./app-0a598c04.js";const t={},o=e(`<h1 id="import" tabindex="-1"><a class="header-anchor" href="#import" aria-hidden="true">#</a> Import</h1><p>Import metadata from source files and create resources into Destination cluster</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token function">import</span> <span class="token parameter variable">--help</span>
</code></pre></div><p>Usage:</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>Command to <span class="token function">import</span> cluster resources  to another cluster.

Usage:
  cctools <span class="token function">import</span> <span class="token punctuation">[</span>flags<span class="token punctuation">]</span>
  cctools <span class="token function">import</span> <span class="token punctuation">[</span>command<span class="token punctuation">]</span>

Aliases:
  import, i

Available Commands:
  topics      Import Topics Info

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> <span class="token function">import</span>

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>
</code></pre></div><h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><p>Source path:</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">import</span><span class="token punctuation">:</span>
  <span class="token key atrule">source</span><span class="token punctuation">:</span> &lt;path<span class="token punctuation">&gt;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>::: info The source path is <code>export.output</code> folder by default. :::</p>`,9),c=[o];function p(i,l){return n(),s("div",null,c)}const u=a(t,[["render",p],["__file","import.html.vue"]]);export{u as default};
