const fs = require("fs");
const path = require("path");

const outDir = path.join(__dirname, "..", "postman");
const collectionPath = path.join(outDir, "e-kinerja-service.postman_collection.json");
const environmentPath = path.join(outDir, "e-kinerja-service.postman_environment.json");

const baseUrl = "{{base_url}}";

const autoSetVariablesScript = `
const json = (() => {
  try { return pm.response.json(); } catch (error) { return null; }
})();

pm.test("Response is valid JSON when body is present", function () {
  if (pm.response.text()) {
    pm.expect(json).to.not.equal(null);
  }
});

if (json) {
  const setIfValue = (name, value) => {
    if (value !== undefined && value !== null && String(value).trim() !== "") {
      pm.environment.set(name, String(value));
    }
  };

  const data = json.data ?? json;
  const first = Array.isArray(data) ? data[0] : data;

  setIfValue("access_token", first?.access_token ?? json.access_token);
  setIfValue("refresh_token", first?.refresh_token ?? json.refresh_token);

  const directKeyMap = {
    id: pm.request.url.path.filter(Boolean)[0]?.replace(/-/g, "_") + "_id",
    role_id: "role_id",
    user_id: "user_id",
    pemda_id: "pemda_id",
    aplikasi_id: "aplikasi_id",
    permintaan_id: "permintaan_id",
    distribusi_id: "distribusi_id",
    distribusi_pelaksana_id: "distribusi_pelaksana_id",
    pelaksana_id: "pelaksana_id",
    programmer_id: "programmer_id",
    laporan_id: "laporan_id",
    verifikasi_id: "verifikasi_id",
    penugasan_id: "penugasan_id",
    penilaian_id: "penilaian_id"
  };

  Object.entries(directKeyMap).forEach(([key, variable]) => {
    setIfValue(variable, first?.[key]);
  });

  const resourceIdMap = {
    "roles": "role_id",
    "users": "user_id",
    "master-aplikasi": "master_aplikasi_id",
    "master-pemda": "master_pemda_id",
    "permintaan": "permintaan_id",
    "distribusi": "distribusi_id",
    "pelaksana": "pelaksana_id",
    "laporan": "laporan_id",
    "verifikasi": "verifikasi_id",
    "penugasan": "penugasan_id",
    "penilaian": "penilaian_id"
  };

  const resource = pm.request.url.path.filter(Boolean)[0];
  setIfValue(resourceIdMap[resource], first?.id);

  if (resource === "master-aplikasi") {
    setIfValue("aplikasi_id", first?.id);
  }

  if (resource === "master-pemda") {
    setIfValue("pemda_id", first?.id);
  }

  if (resource === "roles") {
    setIfValue("role_id", first?.id);
  }

  const visit = (value) => {
    if (!value || typeof value !== "object") return;
    if (Array.isArray(value)) {
      value.forEach(visit);
      return;
    }

    if (value.id && value.name) {
      const name = String(value.name).toLowerCase().replace(/[^a-z0-9]+/g, "_").replace(/^_|_$/g, "");
      if (name) setIfValue(name + "_id", value.id);
    }

    Object.entries(value).forEach(([key, child]) => {
      if (key.endsWith("_id")) setIfValue(key, child);
      if (child && typeof child === "object" && child.id) setIfValue(key.replace(/-/g, "_") + "_id", child.id);
      visit(child);
    });
  };

  visit(data);
}
`;

const defaultHeader = [
  { key: "Accept", value: "application/json" },
  { key: "Content-Type", value: "application/json" }
];

const formHeader = [
  { key: "Accept", value: "application/json" }
];

const raw = (body) => ({
  mode: "raw",
  raw: JSON.stringify(body, null, 2),
  options: { raw: { language: "json" } }
});

const formdata = (fields) => ({
  mode: "formdata",
  formdata: fields.map((field) => ({
    key: field.key,
    value: field.value || "",
    type: field.type || "text",
    src: field.src || []
  }))
});

const auth = { type: "bearer", bearer: [{ key: "token", value: "{{access_token}}", type: "string" }] };
const noauth = { type: "noauth" };

const postmanUrl = (url) => url.replace(/:([A-Za-z0-9_]+)/g, "{{$1}}");

const request = ({ name, method, url, body, header, authOverride }) => ({
  name,
  event: [{ listen: "test", script: { type: "text/javascript", exec: autoSetVariablesScript.trim().split("\n") } }],
  request: {
    method,
    header: header || defaultHeader,
    auth: authOverride,
    url: {
      raw: `${baseUrl}${postmanUrl(url)}`,
      host: ["{{base_url}}"],
      path: url.split("?")[0].split("/").filter(Boolean).map((segment) => segment.startsWith(":") ? `{{${segment.slice(1)}}}` : segment),
      query: (url.split("?")[1] || "").split("&").filter(Boolean).map((query) => {
        const [key, value] = query.split("=");
        return { key, value };
      })
    },
    ...(body ? { body } : {})
  }
});

const folder = (name, items) => ({ name, item: items });

const collection = {
  info: {
    _postman_id: "5aef5952-5966-4dc1-91af-ec748476c251",
    name: "E-Kinerja Service",
    description: "Generated Postman collection for e-kinerja-service. Import the paired environment file, then run Login first so access_token and refresh_token are populated.",
    schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  auth,
  event: [],
  item: [
    folder("Auth", [
      request({ name: "Login", method: "POST", url: "/auth/login", authOverride: noauth, body: raw({ username: "{{username}}", password: "{{password}}" }) }),
      request({ name: "Refresh Token", method: "POST", url: "/auth/refresh", authOverride: noauth, body: raw({ refresh_token: "{{refresh_token}}" }) }),
      request({ name: "Logout", method: "POST", url: "/auth/logout" })
    ]),
    folder("Roles", [
      request({ name: "Get Roles", method: "GET", url: "/roles", authOverride: noauth }),
      request({ name: "Get Role By ID", method: "GET", url: "/roles/:role_id", authOverride: noauth }),
      request({ name: "Create Role", method: "POST", url: "/roles", body: raw({ name: "admin", description: "Administrator" }) }),
      request({ name: "Update Role", method: "PATCH", url: "/roles/:role_id", body: raw({ name: "admin", description: "Administrator updated" }) }),
      request({ name: "Delete Role", method: "DELETE", url: "/roles/:role_id" })
    ]),
    folder("Users", [
      request({ name: "Get Users", method: "GET", url: "/users", authOverride: noauth }),
      request({ name: "Get User By ID", method: "GET", url: "/users/:user_id", authOverride: noauth }),
      request({ name: "Create User", method: "POST", url: "/users", body: raw({ role_id: "{{role_id}}", username: "programmer01", full_name: "Programmer Satu", password: "password123" }) }),
      request({ name: "Update User", method: "PATCH", url: "/users/:user_id", body: raw({ role_id: "{{role_id}}", username: "programmer01", full_name: "Programmer Satu Updated", password: "password123", is_active: true }) }),
      request({ name: "Upload Profile Picture", method: "PATCH", url: "/users/:user_id/profile-picture", header: formHeader, body: formdata([{ key: "profile_picture", type: "file" }]) }),
      request({ name: "Delete User", method: "DELETE", url: "/users/:user_id" })
    ]),
    folder("Master Aplikasi", [
      request({ name: "Get Master Aplikasi", method: "GET", url: "/master-aplikasi" }),
      request({ name: "Get Master Aplikasi By ID", method: "GET", url: "/master-aplikasi/:master_aplikasi_id" }),
      request({ name: "Create Master Aplikasi", method: "POST", url: "/master-aplikasi", body: raw({ name: "E-Kinerja", link: "https://example.com/e-kinerja" }) }),
      request({ name: "Update Master Aplikasi", method: "PUT", url: "/master-aplikasi/:master_aplikasi_id", body: raw({ name: "E-Kinerja Updated", link: "https://example.com/e-kinerja" }) }),
      request({ name: "Upload Logo Aplikasi", method: "PATCH", url: "/master-aplikasi/:master_aplikasi_id/logo", header: formHeader, body: formdata([{ key: "logo", type: "file" }]) }),
      request({ name: "Delete Master Aplikasi", method: "DELETE", url: "/master-aplikasi/:master_aplikasi_id" })
    ]),
    folder("Master Pemda", [
      request({ name: "Get Master Pemda", method: "GET", url: "/master-pemda" }),
      request({ name: "Get Master Pemda By ID", method: "GET", url: "/master-pemda/:master_pemda_id" }),
      request({ name: "Create Master Pemda", method: "POST", url: "/master-pemda", body: raw({ name: "Pemda Contoh", link: "https://example.com/pemda" }) }),
      request({ name: "Update Master Pemda", method: "PUT", url: "/master-pemda/:master_pemda_id", body: raw({ name: "Pemda Contoh Updated", link: "https://example.com/pemda" }) }),
      request({ name: "Upload Logo Pemda", method: "PATCH", url: "/master-pemda/:master_pemda_id/logo", header: formHeader, body: formdata([{ key: "logo", type: "file" }]) }),
      request({ name: "Delete Master Pemda", method: "DELETE", url: "/master-pemda/:master_pemda_id" })
    ]),
    folder("Permintaan", [
      request({ name: "Get Permintaan", method: "GET", url: "/permintaan" }),
      request({ name: "Get Archived Permintaan", method: "GET", url: "/permintaan/archived" }),
      request({ name: "Get Permintaan By ID", method: "GET", url: "/permintaan/:permintaan_id" }),
      request({ name: "Create Permintaan", method: "POST", url: "/permintaan", body: raw({ pemda_id: "{{pemda_id}}", aplikasi_id: "{{aplikasi_id}}", menu: "Dashboard", kondisi_awal: "Belum tersedia", kondisi_diharapkan: "Dashboard tersedia", tanggal_pesanan: "2026-05-17", tanggal_deadline: "2026-06-17" }) }),
      request({ name: "Update Permintaan", method: "PUT", url: "/permintaan/:permintaan_id", body: raw({ pemda_id: "{{pemda_id}}", aplikasi_id: "{{aplikasi_id}}", menu: "Dashboard Updated", kondisi_awal: "Belum tersedia", kondisi_diharapkan: "Dashboard tersedia dan lengkap", tanggal_pesanan: "2026-05-17", tanggal_deadline: "2026-06-17" }) }),
      request({ name: "Update Status Permintaan", method: "PATCH", url: "/permintaan/:permintaan_id/status", body: raw({ status: "diproses" }) }),
      request({ name: "Upload Lampiran Permintaan", method: "PATCH", url: "/permintaan/:permintaan_id/lampiran", header: formHeader, body: formdata([{ key: "lampiran", type: "file" }]) }),
      request({ name: "Delete Permintaan", method: "DELETE", url: "/permintaan/:permintaan_id" })
    ]),
    folder("Distribusi", [
      request({ name: "Get Distribusi", method: "GET", url: "/distribusi?expand=names" }),
      request({ name: "Get Distribusi By ID", method: "GET", url: "/distribusi/:distribusi_id?expand=names" }),
      request({ name: "Create Distribusi", method: "POST", url: "/distribusi", body: raw({ permintaan_id: "{{permintaan_id}}", komentar: "Mohon dikerjakan", programmer_ids: ["{{programmer_id}}"], pelaksana: ["{{programmer_id}}"] }) }),
      request({ name: "Update Distribusi", method: "PUT", url: "/distribusi/:distribusi_id", body: raw({ permintaan_id: "{{permintaan_id}}", komentar: "Mohon dikerjakan dengan prioritas", programmer_ids: ["{{programmer_id}}"], pelaksana: ["{{programmer_id}}"] }) }),
      request({ name: "Create Komentar Distribusi", method: "POST", url: "/distribusi/komentar/:distribusi_id", body: raw({ komentars: "Komentar distribusi" }) }),
      request({ name: "Delete Distribusi", method: "DELETE", url: "/distribusi/:distribusi_id" })
    ]),
    folder("Pelaksana", [
      request({ name: "Get Pelaksana", method: "GET", url: "/pelaksana" }),
      request({ name: "Mark All Read Pelaksana", method: "PATCH", url: "/pelaksana/mark-all-read" }),
      request({ name: "Get Pelaksana By ID", method: "GET", url: "/pelaksana/:pelaksana_id" }),
      request({ name: "Create Pelaksana", method: "POST", url: "/pelaksana", body: raw({ distribusi_id: "{{distribusi_id}}", programmer_id: "{{programmer_id}}" }) }),
      request({ name: "Update Pelaksana", method: "PUT", url: "/pelaksana/:pelaksana_id", body: raw({ distribusi_id: "{{distribusi_id}}", programmer_id: "{{programmer_id}}" }) }),
      request({ name: "Delete Pelaksana", method: "DELETE", url: "/pelaksana/:pelaksana_id" })
    ]),
    folder("Laporan", [
      request({ name: "Get Laporan", method: "GET", url: "/laporan" }),
      request({ name: "Get History Laporan", method: "GET", url: "/laporan/history" }),
      request({ name: "Get Laporan By ID", method: "GET", url: "/laporan/:laporan_id" }),
      request({ name: "Create Laporan", method: "POST", url: "/laporan", header: formHeader, body: formdata([{ key: "permintaan_id", value: "{{permintaan_id}}" }, { key: "penugasan_id", value: "{{penugasan_id}}" }, { key: "laporan_progress", value: "Progress pekerjaan" }, { key: "status", value: "sedang_berjalan" }, { key: "lampiran", type: "file" }]) }),
      request({ name: "Create Verif From Laporan", method: "POST", url: "/laporan/verif/:laporan_id", body: raw({ laporan_id: "{{laporan_id}}" }) }),
      request({ name: "Update Laporan", method: "PUT", url: "/laporan/:laporan_id", body: raw({ permintaan_id: "{{permintaan_id}}", laporan_progress: "Progress pekerjaan updated", status: "selesai", verifikasi_id: "{{verifikasi_id}}", status_verified: "pending", is_submitted_to_verified: true }) }),
      request({ name: "Upload Lampiran Laporan", method: "PATCH", url: "/laporan/:laporan_id/lampiran", header: formHeader, body: formdata([{ key: "lampiran", type: "file" }]) }),
      request({ name: "Create Komentar Laporan", method: "POST", url: "/laporan/komentar/:laporan_id", body: raw({ komentar: "Komentar laporan" }) }),
      request({ name: "Delete Laporan", method: "DELETE", url: "/laporan/:laporan_id" })
    ]),
    folder("Verifikasi", [
      request({ name: "Get Verifikasi", method: "GET", url: "/verifikasi" }),
      request({ name: "Get Verifikasi By ID", method: "GET", url: "/verifikasi/:verifikasi_id" }),
      request({ name: "Create Verifikasi", method: "POST", url: "/verifikasi", body: raw({ laporan_id: "{{laporan_id}}", komentar: "Sudah diperiksa", status_verified: "approved" }) }),
      request({ name: "Update Verifikasi", method: "PUT", url: "/verifikasi/:verifikasi_id", body: raw({ laporan_id: "{{laporan_id}}", komentar: "Perlu revisi", status_verified: "revision" }) }),
      request({ name: "Delete Verifikasi", method: "DELETE", url: "/verifikasi/:verifikasi_id" })
    ]),
    folder("Penugasan", [
      request({ name: "Get Penugasan", method: "GET", url: "/penugasan?pelaksana_id={{pelaksana_id}}&distribusi_id={{distribusi_id}}" }),
      request({ name: "Get Penugasan By ID", method: "GET", url: "/penugasan/:penugasan_id" }),
      request({ name: "Create Penugasan", method: "POST", url: "/penugasan", body: raw({ distribusi_pelaksana_id: "{{distribusi_pelaksana_id}}", judul: "Implementasi fitur", deskripsi: "Mengerjakan fitur sesuai permintaan", deadline: "2026-06-17", prioritas: "medium", estimasi_hari: 5, urutan: 1 }) }),
      request({ name: "Update Penugasan", method: "PUT", url: "/penugasan/:penugasan_id", body: raw({ judul: "Implementasi fitur updated", deskripsi: "Mengerjakan fitur sesuai permintaan", deadline: "2026-06-17", prioritas: "high", estimasi_hari: 3, urutan: 1 }) }),
      request({ name: "Update Status Penugasan", method: "PATCH", url: "/penugasan/:penugasan_id/status", body: raw({ status: "sedang_berjalan" }) }),
      request({ name: "Reassign Penugasan", method: "PATCH", url: "/penugasan/:penugasan_id/reassign", body: raw({ distribusi_pelaksana_id: "{{distribusi_pelaksana_id}}" }) }),
      request({ name: "Delete Penugasan", method: "DELETE", url: "/penugasan/:penugasan_id" })
    ]),
    folder("Penilaian", [
      request({ name: "Get Penilaian", method: "GET", url: "/penilaian" }),
      request({ name: "Get Penilaian By Distribusi", method: "GET", url: "/penilaian/:distribusi_id" }),
      request({ name: "Create Penilaian", method: "POST", url: "/penilaian", body: raw({ distribusi_id: "{{distribusi_id}}", tingkat_keberhasilan: 90, komentar: "Pekerjaan baik", tanggal_selesai: "2026-06-17" }) }),
      request({ name: "Update Penilaian", method: "PATCH", url: "/penilaian/:penilaian_id", body: raw({ tingkat_keberhasilan: 95, komentar: "Pekerjaan sangat baik", tanggal_selesai: "2026-06-17" }) })
    ]),
    folder("Dashboard & Activity", [
      request({ name: "Get Superadmin Dashboard", method: "GET", url: "/superadmin-dashboard" }),
      request({ name: "Get All Activity", method: "GET", url: "/all-activity" })
    ])
  ]
};

const environment = {
  id: "43afcd71-3d4f-4e5b-87a2-a9b4c5a03402",
  name: "E-Kinerja Service Local",
  values: [
    ["base_url", "http://localhost:8082"],
    ["username", "superadmin"],
    ["password", "password123"],
    ["access_token", ""],
    ["refresh_token", ""],
    ["role_id", ""],
    ["user_id", ""],
    ["programmer_id", ""],
    ["pemda_id", ""],
    ["master_pemda_id", ""],
    ["aplikasi_id", ""],
    ["master_aplikasi_id", ""],
    ["permintaan_id", ""],
    ["distribusi_id", ""],
    ["distribusi_pelaksana_id", ""],
    ["pelaksana_id", ""],
    ["laporan_id", ""],
    ["verifikasi_id", ""],
    ["penugasan_id", ""],
    ["penilaian_id", ""]
  ].map(([key, value]) => ({ key, value, type: "default", enabled: true })),
  _postman_variable_scope: "environment",
  _postman_exported_using: "Codex"
};

fs.mkdirSync(outDir, { recursive: true });
fs.writeFileSync(collectionPath, JSON.stringify(collection, null, 2) + "\n");
fs.writeFileSync(environmentPath, JSON.stringify(environment, null, 2) + "\n");

console.log(`Wrote ${collectionPath}`);
console.log(`Wrote ${environmentPath}`);
