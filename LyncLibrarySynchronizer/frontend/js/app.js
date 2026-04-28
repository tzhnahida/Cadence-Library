// PCB Library Assistant — Frontend
// Uses Wails v3 generated bindings in ../bindings/pcb-library/

// Accept both possible binding paths (wails3 dev vs production)
let AnalyzeLCSC, SaveToDatabase, ClearHistory;
let GetConfig, UpdateConfig, GetDBStats;

async function loadBindings() {
    try {
        const appMod = await import('../bindings/pcb-library/appservice.js');
        const cfgMod = await import('../bindings/pcb-library/configservice.js');
        AnalyzeLCSC = appMod.AnalyzeLCSC;
        SaveToDatabase = appMod.SaveToDatabase;
        ClearHistory = appMod.ClearHistory;
        GetConfig = cfgMod.GetConfig;
        UpdateConfig = cfgMod.UpdateConfig;
        GetDBStats = cfgMod.GetDBStats;
        return true;
    } catch (e) {
        console.error('Bindings load failed:', e);
        return false;
    }
}

// --- State ---
let currentResult = null;

// --- DOM refs ---
const $ = (s) => document.querySelector(s);

let urlInput, analyzeBtn, saveBtn, saveFeedback;
let loadingSection, resultSection, errorSection, errorBanner, inputError;
let fieldsBody, tableSelect, lcscIdDisplay, tagDisplay, tagText;
let statusBar, statsContainer, configPanel;

function cacheDom() {
    urlInput = $('#url-input');
    analyzeBtn = $('#analyze-btn');
    saveBtn = $('#save-btn');
    saveFeedback = $('#save-feedback');
    loadingSection = $('#loading-section');
    resultSection = $('#result-section');
    errorSection = $('#error-section');
    errorBanner = $('#error-banner');
    inputError = $('#input-error');
    fieldsBody = $('#fields-body');
    tableSelect = $('#table-select');
    lcscIdDisplay = $('#lcsc-id-display');
    tagDisplay = $('#tag-display');
    tagText = $('#tag-text');
    statusBar = $('#status-bar');
    statsContainer = $('#stats-container');
    configPanel = $('#config-panel');
}

// --- Init ---
async function init() {
    cacheDom();

    // Show "loading" until bindings resolve
    if (statusBar) statusBar.textContent = '加载中...';

    const ok = await loadBindings();
    if (!ok) {
        if (statusBar) statusBar.textContent = '错误: 无法加载 Wails 运行时';
        if (errorSection) { errorSection.classList.remove('hidden'); errorBanner.textContent = 'Wails 运行时加载失败，请重启应用'; }
        return;
    }

    setupEvents();
    refreshStats();
    loadConfig();
    if (statusBar) statusBar.textContent = '就绪';
}

function setupEvents() {
    if (analyzeBtn) analyzeBtn.addEventListener('click', handleAnalyze);
    if (saveBtn) saveBtn.addEventListener('click', handleSave);
    const clearBtn = $('#clear-btn');
    if (clearBtn) clearBtn.addEventListener('click', handleClearHistory);
    const configToggleBtn = $('#config-toggle');
    if (configToggleBtn) configToggleBtn.addEventListener('click', toggleConfig);
    const cfgSaveBtn = $('#cfg-save');
    if (cfgSaveBtn) cfgSaveBtn.addEventListener('click', saveConfig);
    if (urlInput) urlInput.addEventListener('keydown', (e) => {
        if (e.key === 'Enter') handleAnalyze();
    });
    if (tableSelect) tableSelect.addEventListener('change', () => {
        if (currentResult) renderFields(currentResult);
    });
}

// --- Analyze ---
async function handleAnalyze() {
    if (!urlInput) return;
    const url = urlInput.value.trim();
    if (!url) {
        showInputError('请输入 URL');
        return;
    }
    showInputError(null);
    setLoading(true);
    hideError();

    try {
        currentResult = await AnalyzeLCSC(url);
        renderResult(currentResult);
        hideError();
        if (statusBar) statusBar.textContent = `分析完成: ${currentResult.lcsc_id}`;
    } catch (err) {
        currentResult = null;
        showError('分析失败: ' + err);
        if (statusBar) statusBar.textContent = '分析失败';
    }
    setLoading(false);
}

function renderResult(result) {
    if (lcscIdDisplay) lcscIdDisplay.textContent = 'LCSC ID: ' + (result.lcsc_id || '-');

    if (tableSelect) {
        tableSelect.innerHTML = '';
        if (result.table_names) {
            result.table_names.forEach(name => {
                const opt = document.createElement('option');
                opt.value = name;
                opt.textContent = name;
                if (name === result.table_name) opt.selected = true;
                tableSelect.appendChild(opt);
            });
        }
    }

    renderFields(result);
    if (resultSection) resultSection.classList.remove('hidden');
    if (tagDisplay) tagDisplay.classList.add('hidden');
    if (saveFeedback) saveFeedback.textContent = '';
}

function renderFields(result) {
    if (!fieldsBody || !result.fields) return;

    fieldsBody.innerHTML = '';
    const keys = Object.keys(result.fields);
    keys.sort((a, b) => {
        if (a === 'Part_ID') return -1;
        if (b === 'Part_ID') return 1;
        if (a === 'Category') return -1;
        if (b === 'Category') return 1;
        if (a === 'Symbol_Name') return -1;
        if (b === 'Symbol_Name') return 1;
        return a.localeCompare(b);
    });

    keys.forEach(key => {
        const tr = document.createElement('tr');
        const tdKey = document.createElement('td');
        tdKey.textContent = key;
        const tdVal = document.createElement('td');
        const input = document.createElement('input');
        input.type = 'text';
        input.value = result.fields[key] || '';
        input.dataset.field = key;
        input.addEventListener('input', () => {
            result.fields[key] = input.value;
        });
        tdVal.appendChild(input);
        tr.appendChild(tdKey);
        tr.appendChild(tdVal);
        fieldsBody.appendChild(tr);
    });
}

// --- Save ---
async function handleSave() {
    if (!currentResult) return;

    if (fieldsBody) {
        fieldsBody.querySelectorAll('input').forEach(input => {
            currentResult.fields[input.dataset.field] = input.value;
        });
    }
    if (tableSelect) currentResult.table_name = tableSelect.value;

    if (saveBtn) saveBtn.disabled = true;
    if (saveFeedback) { saveFeedback.textContent = '写入中...'; saveFeedback.className = ''; }

    try {
        const result = await SaveToDatabase(currentResult);
        if (saveFeedback) { saveFeedback.textContent = `成功! Part_ID: ${result.part_id}`; saveFeedback.className = 'text-success'; }
        if (tagText) tagText.textContent = result.system_tag || '';
        if (tagDisplay) tagDisplay.classList.remove('hidden');
        if (statusBar) statusBar.textContent = `已保存 Part_ID: ${result.part_id}`;
        refreshStats();
    } catch (err) {
        if (saveFeedback) { saveFeedback.textContent = '写入失败: ' + err; saveFeedback.className = 'text-error'; }
    }
    if (saveBtn) saveBtn.disabled = false;
}

// --- Clear History ---
async function handleClearHistory() {
    try {
        await ClearHistory();
        if (statusBar) statusBar.textContent = '对话历史已清除';
    } catch (err) {
        if (statusBar) statusBar.textContent = '清除失败: ' + err;
    }
}

// --- Config ---
function toggleConfig() {
    if (configPanel) configPanel.classList.toggle('hidden');
}

async function loadConfig() {
    try {
        const cfg = await GetConfig();
        setVal('#cfg-apikey', cfg.ApiKey);
        setVal('#cfg-baseurl', cfg.BaseUrl);
        setVal('#cfg-model', cfg.AiModel);
        setVal('#cfg-dbfile', cfg.DbFile);
        setVal('#cfg-systemtag', cfg.SystemTag);
    } catch (err) {
        console.error('Load config failed:', err);
    }
}

function setVal(sel, val) {
    const el = document.querySelector(sel);
    if (el) el.value = val || '';
}

async function saveConfig() {
    const cfgBtn = $('#cfg-save');
    const fb = $('#cfg-feedback');
    if (cfgBtn) cfgBtn.disabled = true;
    if (fb) { fb.textContent = '保存中...'; fb.className = ''; }

    try {
        await UpdateConfig({
            ApiKey: getVal('#cfg-apikey'),
            BaseUrl: getVal('#cfg-baseurl'),
            AiModel: getVal('#cfg-model'),
            DbFile: getVal('#cfg-dbfile'),
            SystemTag: getVal('#cfg-systemtag'),
        });
        if (fb) { fb.textContent = '已保存'; fb.className = 'text-success'; }
    } catch (err) {
        if (fb) { fb.textContent = '保存失败: ' + err; fb.className = 'text-error'; }
    }
    if (cfgBtn) cfgBtn.disabled = false;
}

function getVal(sel) {
    const el = document.querySelector(sel);
    return el ? el.value : '';
}

// --- Stats ---
async function refreshStats() {
    if (!statsContainer) return;
    try {
        const stats = await GetDBStats();
        if (!stats || stats.length === 0) {
            statsContainer.innerHTML = '<span class="loading-tip">无数据</span>';
            return;
        }
        let html = '<table class="stats-table"><tr><th>表名</th><th>记录</th></tr>';
        stats.forEach(s => {
            html += `<tr><td>${s.name}</td><td>${s.records}</td></tr>`;
        });
        html += '</table>';
        statsContainer.innerHTML = html;
    } catch (err) {
        statsContainer.innerHTML = '<span class="loading-tip">无法连接数据库</span>';
    }
}

// --- UI helpers ---
function setLoading(loading) {
    if (loadingSection) loadingSection.classList.toggle('hidden', !loading);
    if (analyzeBtn) analyzeBtn.disabled = loading;
    if (loading && resultSection) resultSection.classList.add('hidden');
}

function showInputError(msg) {
    if (inputError) {
        inputError.textContent = msg || '';
        inputError.classList.toggle('hidden', !msg);
    }
}

function showError(msg) {
    if (errorBanner) errorBanner.textContent = msg;
    if (errorSection) errorSection.classList.remove('hidden');
    if (resultSection) resultSection.classList.add('hidden');
}

function hideError() {
    if (errorSection) errorSection.classList.add('hidden');
}

// --- Start ---
init();
