-- ============================================================
-- SEED DUMMY DATA
-- Aplikasi Stok Barang Wifi Logistik Berbasis Web
-- PT EKA MAS REPUBLIK (MyRepublic Solo)
-- Urutan Insert: categories → suppliers → products → stock_ins → stock_outs
-- ============================================================

-- ------------------------------------------------------------
-- 1. CATEGORIES (20 rows)
-- Konteks: Kategori barang jaringan/wifi/logistik ISP
-- ------------------------------------------------------------
INSERT INTO categories (created_at, updated_at, created_by_id, updated_by_id, name, description) VALUES
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Router & Access Point',    'Perangkat router dan access point untuk distribusi sinyal WiFi'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Kabel & Fiber Optik',      'Kabel UTP, FTP, dan kabel fiber optik untuk instalasi jaringan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Splitter & ODP',           'Optical Distribution Point dan splitter fiber optik'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'ONT & CPE',                'Optical Network Terminal dan perangkat Customer Premises Equipment'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Switch & Hub',             'Network switch dan hub untuk distribusi jaringan LAN'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Tools Instalasi',          'Peralatan teknis untuk instalasi jaringan di lapangan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Konektor & Adaptor',       'Konektor RJ45, SC, LC, dan adaptor jaringan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Tiang & Bracket',          'Tiang antena, bracket wall mount, dan aksesoris pemasangan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Power Supply & UPS',       'Adaptor daya, PoE injector, dan UPS untuk perangkat jaringan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Closure & Joint',          'Closure fiber optik dan joint protection sleeve'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Konduit & Pipa',           'Konduit kabel, pipa PVC, dan klem untuk proteksi kabel'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Perangkat CCTV',           'Kamera CCTV dan aksesoris sistem pengawasan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'OTB & Patch Panel',        'Optical Terminal Box dan patch panel untuk manajemen kabel'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Patchcord & Pigtail',      'Kabel patchcord fiber dan pigtail untuk terminasi fiber'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Aksesoris Lapangan',       'Tie wrap, label kabel, dan aksesoris instalasi lapangan'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'APD & Safety',             'Alat Pelindung Diri dan peralatan keselamatan kerja teknisi'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'ATK & Administrasi',       'Alat tulis kantor dan kebutuhan administrasi operasional'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Perangkat Monitoring',     'Alat ukur dan monitoring jaringan, seperti OTDR dan LAN tester'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Server & Rack',            'Server, rack cabinet, dan aksesoris ruang NOC'),
('2024-01-02 08:00:00', '2024-01-02 08:00:00', 1, 1, 'Spare Part Perangkat',     'Suku cadang dan komponen pengganti perangkat jaringan');


-- ------------------------------------------------------------
-- 2. SUPPLIERS (20 rows)
-- Konteks: Supplier perangkat jaringan/ISP wilayah Solo & nasional
-- ------------------------------------------------------------
INSERT INTO suppliers (created_at, updated_at, created_by_id, updated_by_id, code, name, contact, phone, email, address, status) VALUES
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-001', 'PT Solusi Jaringan Indonesia',   'Andi Kurniawan',    '021-7788001', 'andi@sji.co.id',             'Jl. Gatot Subroto Kav.9, Jakarta Selatan',      'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-002', 'CV Mitra Telekomunikasi Solo',   'Bambang Susilo',    '0271-772002', 'bambang@mitrasolo.com',      'Jl. Slamet Riyadi No.88, Solo',                 'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-003', 'PT Indotech Distribusi',         'Cahya Nugroho',     '022-6543003', 'cahya@indotech.co.id',       'Jl. Asia Afrika No.45, Bandung',                'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-004', 'UD Fiber Optik Jaya',            'Dwi Santoso',       '0271-774004', 'dwi@fojaya.com',             'Jl. Veteran No.12, Sukoharjo',                  'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-005', 'PT Mikrotik Authorized Dist.',   'Eka Prasetya',      '031-5675005', 'eka@mikrotik-dist.co.id',    'Jl. Raya Gubeng No.70, Surabaya',               'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-006', 'CV Cahaya Network Solo',         'Fajar Wibowo',      '0271-776006', 'fajar@cahyanet.com',         'Jl. Monginsidi No.34, Solo',                    'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-007', 'PT Global Connectivity',         'Gita Permatasari',  '021-7790007', 'gita@globalconn.co.id',      'Jl. TB Simatupang No.12, Jakarta Selatan',      'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-008', 'UD Toko Jaringan Klaten',        'Hendra Setiawan',   '0272-778008', 'hendra@tokojaring.com',      'Jl. Pemuda No.5, Klaten',                       'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-009', 'PT Nexwave Teknologi',           'Indra Laksono',     '024-7679009', 'indra@nexwave.co.id',        'Jl. Pandanaran No.55, Semarang',                'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-010', 'CV Sarana ISP Jateng',           'Joko Purnomo',      '0271-770010', 'joko@saranaispjt.com',       'Jl. Adi Sucipto No.110, Solo',                  'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-011', 'PT Ubiquiti Indonesia',          'Krisna Adiputra',   '021-7791011', 'krisna@ubnt-indo.co.id',     'Jl. Casablanca No.88, Jakarta Selatan',         'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-012', 'CV Optika Fiber Solo',           'Lina Marliana',     '0271-772012', 'lina@optikafibersolo.com',   'Jl. Honggowongso No.22, Solo',                  'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-013', 'PT Telko Supplies Nasional',     'Maulana Yusuf',     '022-6541013', 'maulana@telkosupply.co.id',  'Jl. Braga No.30, Bandung',                      'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-014', 'UD Peralatan Network Boyolali',  'Novi Rahayu',       '0276-774014', 'novi@netboyolali.com',       'Jl. Pandanaran No.9, Boyolali',                 'inactive'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-015', 'PT TP-Link Authorized Solo',     'Oscar Pratama',     '0271-775015', 'oscar@tplink-solo.co.id',    'Jl. Dr. Radjiman No.66, Solo',                  'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-016', 'CV Huawei Reseller Jateng',      'Putri Handayani',   '024-7676016', 'putri@huawei-jt.com',        'Jl. MT Haryono No.8, Semarang',                 'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-017', 'PT Jaringan Nusantara',          'Rizky Firmansyah',  '021-7797017', 'rizky@jarnusa.co.id',        'Jl. HR Rasuna Said No.10, Jakarta',             'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-018', 'CV Mega Tech Solo',              'Sandi Wijaya',      '0271-778018', 'sandi@megatechsolo.com',     'Jl. Urip Sumoharjo No.40, Solo',                'active'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-019', 'UD Elektro Prima Wonogiri',      'Tania Kusuma',      '0273-779019', 'tania@elektroprima.com',     'Jl. A. Yani No.15, Wonogiri',                   'inactive'),
('2024-01-03 08:00:00', '2024-01-03 08:00:00', 1, 1, 'SUP-020', 'PT Cisco Authorized Partner',    'Udin Syahrudin',    '021-7792020', 'udin@cisco-partner.co.id',   'Jl. Jend. Sudirman Kav.25, Jakarta Pusat',      'active');


-- ------------------------------------------------------------
-- 3. PRODUCTS (20 rows)
-- Konteks: Barang stok operasional WiFi / ISP PT EKA MAS REPUBLIK
-- ------------------------------------------------------------
INSERT INTO products (created_at, updated_at, created_by_id, updated_by_id, code, name, category_id, supplier_id, description, stock, min_stock, unit, price) VALUES
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-001', 'MikroTik hAP ac² RB952Ui-5ac2nD',  1,  5,  'Dual-band home AP 2.4GHz/5GHz 802.11ac, 5 port Ethernet, cocok untuk pelanggan rumahan MyRepublic', 35,  10, 'Unit',    850000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-002', 'Kabel Fiber Optik FTTH G.657A 1km', 2,  4,  'Kabel fiber drop FTTH single mode G.657A, tahan tekuk, untuk instalasi pelanggan residensial',       20,  5,  'Roll',    650000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-003', 'ODP 8 Core Outdoor',               3,  12, 'Optical Distribution Point 8 core outdoor, dilengkapi splitter 1:8, untuk distribusi fiber ke rumah',  15,  4,  'Unit',    580000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-004', 'ONT Huawei EG8145V5',              4,  16, 'ONT 4 port LAN + WiFi dual band, GPON, digunakan sebagai modem pelanggan MyRepublic paket Giga',       40,  10, 'Unit',    750000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-005', 'Switch TP-Link TL-SG1008D 8-Port',  5,  15, 'Switch unmanaged 8 port Gigabit, untuk distribusi LAN di gedung/cluster pelanggan bisnis',            20,  5,  'Unit',    320000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-006', 'Tang Krimping RJ45 + Stripper',     6,  6,  'Set tang krimping kabel UTP RJ45 dan stripper kabel, kebutuhan teknisi lapangan',                      15,  5,  'Set',     120000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-007', 'Konektor SC/APC (100pcs)',           7,  12, 'Konektor SC/APC untuk terminasi fiber optik, kompatibel dengan splitter dan ODP MyRepublic',            50,  10, 'Box',     285000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-008', 'Tiang Rooftop Antena 3m',           8,  18, 'Tiang besi galvanis tinggi 3 meter untuk pemasangan antena/AP di rooftop pelanggan',                    25,  8,  'Batang',  195000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-009', 'PoE Injector 48V Passive',          9,  7,  'Power over Ethernet injector 48V passive untuk supply daya perangkat AP outdoor via kabel UTP',         30,  8,  'Unit',    115000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-010', 'Closure Fiber Optik 24 Core',       10, 4,  'Dome closure fiber optik kapasitas 24 core untuk joint kabel FO aerial dan underground',                18,  5,  'Unit',    420000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-011', 'Konduit PVC 20mm x 4m',             11, 8,  'Konduit/pipa PVC diameter 20mm panjang 4 meter, untuk proteksi kabel FO di dinding pelanggan',           200, 50, 'Batang',  22000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-012', 'Kabel UTP Cat6 305m Belden',        2,  9,  'Kabel UTP Cat6 unshielded 305 meter per rol, merk Belden, untuk instalasi LAN indoor',                  10,  3,  'Roll',    1250000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-013', 'OTB 12 Core Outdoor',               13, 12, 'Optical Terminal Box 12 core untuk terminasi kabel FO, dilengkapi tray splice, mounting pole/wall',       12,  3,  'Unit',    310000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-014', 'Pigtail SC/APC 1.5m (12pcs)',       14, 12, 'Pigtail fiber SC/APC panjang 1.5m untuk splicing di ODP/OTB, isi 12 pcs per pack',                      30,  8,  'Pack',    145000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-015', 'Tie Wrap / Cable Tie 30cm (100pcs)',15, 6,  'Cable tie / tie wrap ukuran 30cm warna hitam, isi 100 pcs, untuk merapikan instalasi kabel lapangan',     150, 30, 'Pack',    18000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-016', 'Helm Safety Teknisi',               16, 10, 'Helm keselamatan kerja untuk teknisi lapangan, SNI, digunakan saat instalasi di tiang/rooftop',           20,  5,  'Unit',    85000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-017', 'Kertas HVS A4 80gsm (1 Rim)',       17, 2,  'Kertas HVS A4 80gsm 500 lembar, untuk kebutuhan administrasi dan cetak dokumen instalasi',               80,  20, 'Rim',     55000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-018', 'OTDR Grandway FHO5000-D26',         18, 9,  'Optical Time Domain Reflectometer untuk pengukuran dan troubleshooting jaringan fiber optik di lapangan',  3,   1,  'Unit',    18500000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-019', 'Rack Server 12U Wallmount',         19, 7,  'Rack cabinet wallmount 12U untuk NOC / ruang server cabang Solo, dilengkapi fan dan kunci',                2,   1,  'Unit',    3750000),
('2024-01-05 08:00:00', '2024-01-05 08:00:00', 1, 1, 'PRD-020', 'Adaptor/Charger ONT 12V 1A',        20, 15, 'Adaptor daya pengganti untuk ONT pelanggan, output 12V 1A, spare part klaim garansi & penggantian',        60,  15, 'Unit',    35000);


-- ------------------------------------------------------------
-- 4. STOCK IN (20 rows)
-- Konteks: Penerimaan barang gudang PT EKA MAS REPUBLIK Solo
-- ------------------------------------------------------------
INSERT INTO stock_ins (created_at, updated_at, created_by_id, updated_by_id, code, date, product_id, supplier_id, quantity, notes) VALUES
('2024-02-01 08:30:00', '2024-02-01 08:30:00', 1, 1, 'SI-20240201-001', '2024-02-01', 1,  5,  20,  'Pengadaan router MikroTik untuk stok instalasi pelanggan baru paket MyRepublic bulan Februari'),
('2024-02-03 09:00:00', '2024-02-03 09:00:00', 1, 1, 'SI-20240203-001', '2024-02-03', 2,  4,  10,  'Restock kabel fiber FTTH untuk proyek perluasan coverage area Solo Barat'),
('2024-02-05 10:00:00', '2024-02-05 10:00:00', 1, 1, 'SI-20240205-001', '2024-02-05', 3,  12, 8,   'Pengadaan ODP 8 core untuk perluasan distribusi jaringan wilayah Banjarsari'),
('2024-02-07 08:00:00', '2024-02-07 08:00:00', 1, 1, 'SI-20240207-001', '2024-02-07', 4,  16, 25,  'Penerimaan ONT Huawei untuk paket Giga pelanggan residensial dan SOHO'),
('2024-02-09 09:30:00', '2024-02-09 09:30:00', 1, 1, 'SI-20240209-001', '2024-02-09', 5,  15, 10,  'Pembelian switch 8 port untuk distribusi LAN di gedung pelanggan korporat Solo'),
('2024-02-12 10:00:00', '2024-02-12 10:00:00', 1, 1, 'SI-20240212-001', '2024-02-12', 6,  6,  8,   'Pengadaan tang krimping untuk teknisi baru tim instalasi lapangan'),
('2024-02-14 08:30:00', '2024-02-14 08:30:00', 1, 1, 'SI-20240214-001', '2024-02-14', 7,  12, 30,  'Restock konektor SC/APC, stok hampir habis setelah proyek cluster Colomadu'),
('2024-02-16 09:00:00', '2024-02-16 09:00:00', 1, 1, 'SI-20240216-001', '2024-02-16', 8,  18, 15,  'Pembelian tiang rooftop untuk instalasi AP outdoor pelanggan bisnis'),
('2024-02-19 10:30:00', '2024-02-19 10:30:00', 1, 1, 'SI-20240219-001', '2024-02-19', 9,  7,  20,  'Pengadaan PoE injector untuk mendukung instalasi AP outdoor area Laweyan'),
('2024-02-21 08:00:00', '2024-02-21 08:00:00', 1, 1, 'SI-20240221-001', '2024-02-21', 10, 4,  10,  'Penerimaan closure fiber 24 core untuk penyambungan kabel aerial di jalan utama'),
('2024-03-01 09:00:00', '2024-03-01 09:00:00', 1, 1, 'SI-20240301-001', '2024-03-01', 11, 8,  100, 'Restock konduit PVC untuk keperluan instalasi indoor pelanggan bulan Maret'),
('2024-03-04 10:00:00', '2024-03-04 10:00:00', 1, 1, 'SI-20240304-001', '2024-03-04', 12, 9,  5,   'Pembelian kabel UTP Cat6 Belden untuk kebutuhan LAN proyek gedung kampus'),
('2024-03-06 08:30:00', '2024-03-06 08:30:00', 1, 1, 'SI-20240306-001', '2024-03-06', 13, 12, 6,   'Pengadaan OTB 12 core untuk terminasi kabel FO di node distribusi baru'),
('2024-03-08 09:30:00', '2024-03-08 09:30:00', 1, 1, 'SI-20240308-001', '2024-03-08', 14, 12, 20,  'Restock pigtail SC/APC untuk kebutuhan splicing tim FO lapangan'),
('2024-03-11 10:00:00', '2024-03-11 10:00:00', 1, 1, 'SI-20240311-001', '2024-03-11', 15, 6,  80,  'Pembelian cable tie untuk rapi-kabel instalasi batch Maret'),
('2024-03-13 08:00:00', '2024-03-13 08:00:00', 1, 1, 'SI-20240313-001', '2024-03-13', 16, 10, 10,  'Pengadaan helm safety untuk teknisi lapangan baru tim instalasi MyRepublic Solo'),
('2024-03-15 09:00:00', '2024-03-15 09:00:00', 1, 1, 'SI-20240315-001', '2024-03-15', 17, 2,  50,  'Pembelian kertas HVS untuk operasional administrasi kantor Solo Q1 2024'),
('2024-03-18 10:30:00', '2024-03-18 10:30:00', 1, 1, 'SI-20240318-001', '2024-03-18', 18, 9,  1,   'Pembelian OTDR unit cadangan untuk kebutuhan troubleshooting tim FO'),
('2024-03-20 08:30:00', '2024-03-20 08:30:00', 1, 1, 'SI-20240320-001', '2024-03-20', 19, 7,  1,   'Pengadaan rack server wallmount 12U untuk upgrade ruang NOC cabang Solo'),
('2024-03-22 09:00:00', '2024-03-22 09:00:00', 1, 1, 'SI-20240322-001', '2024-03-22', 20, 15, 40,  'Restock adaptor ONT 12V untuk keperluan klaim garansi dan penggantian di lapangan');


-- ------------------------------------------------------------
-- 5. STOCK OUT (20 rows)
-- Konteks: Pengeluaran barang untuk instalasi & operasional lapangan
-- ------------------------------------------------------------
INSERT INTO stock_outs (created_at, updated_at, created_by_id, updated_by_id, code, date, product_id, destination, quantity, notes) VALUES
('2024-02-06 08:00:00', '2024-02-06 08:00:00', 1, 1, 'SO-20240206-001', '2024-02-06', 1,  'Tim Instalasi - Area Laweyan Solo',                   5,  'Distribusi router MikroTik untuk instalasi 5 pelanggan baru paket 50Mbps area Laweyan'),
('2024-02-08 09:00:00', '2024-02-08 09:00:00', 1, 1, 'SO-20240208-001', '2024-02-08', 2,  'Tim FO - Proyek Coverage Banjarsari',                 3,  'Pengeluaran kabel fiber FTTH untuk penarikan kabel ke kluster baru Banjarsari'),
('2024-02-10 10:00:00', '2024-02-10 10:00:00', 1, 1, 'SO-20240210-001', '2024-02-10', 4,  'Tim Instalasi - Area Jebres Solo',                    10, 'ONT Huawei untuk instalasi pelanggan paket Giga di perumahan Jebres'),
('2024-02-13 08:30:00', '2024-02-13 08:30:00', 1, 1, 'SO-20240213-001', '2024-02-13', 3,  'Node Distribusi Colomadu',                            3,  'Pemasangan ODP baru untuk ekspansi jaringan ke kluster perumahan Colomadu'),
('2024-02-15 09:30:00', '2024-02-15 09:30:00', 1, 1, 'SO-20240215-001', '2024-02-15', 9,  'Tim Instalasi - AP Outdoor Serengan',                 8,  'PoE injector untuk supply daya AP outdoor di titik-titik distribusi area Serengan'),
('2024-02-17 10:00:00', '2024-02-17 10:00:00', 1, 1, 'SO-20240217-001', '2024-02-17', 7,  'Tim FO - Terminasi Kabel Solo Barat',                 15, 'Konektor SC/APC untuk terminasi fiber di ODP dan OTB wilayah Solo Barat'),
('2024-02-20 08:00:00', '2024-02-20 08:00:00', 1, 1, 'SO-20240220-001', '2024-02-20', 11, 'Tim Instalasi Indoor - Berbagai Lokasi',              50, 'Konduit PVC untuk proteksi kabel instalasi indoor pelanggan bulan Februari'),
('2024-02-22 09:30:00', '2024-02-22 09:30:00', 1, 1, 'SO-20240222-001', '2024-02-22', 8,  'Tim Instalasi - Pelanggan Bisnis Pasar Kliwon',       5,  'Tiang rooftop untuk pemasangan AP outdoor pelanggan bisnis kawasan Pasar Kliwon'),
('2024-02-24 10:00:00', '2024-02-24 10:00:00', 1, 1, 'SO-20240224-001', '2024-02-24', 5,  'Pelanggan Korporat - RS Kasih Ibu Solo',              2,  'Switch 8 port untuk distribusi LAN di ruang rawat inap RS Kasih Ibu'),
('2024-02-27 08:30:00', '2024-02-27 08:30:00', 1, 1, 'SO-20240227-001', '2024-02-27', 10, 'Tim FO - Penyambungan Kabel Jl. Adi Sucipto',         4,  'Closure fiber 24 core untuk joint kabel aerial yang putus akibat pekerjaan PLN'),
('2024-03-02 09:00:00', '2024-03-02 09:00:00', 1, 1, 'SO-20240302-001', '2024-03-02', 14, 'Tim Splicing FO - Area Mojosongo',                    10, 'Pigtail SC/APC untuk splicing terminasi fiber di ODP area Mojosongo'),
('2024-03-05 10:00:00', '2024-03-05 10:00:00', 1, 1, 'SO-20240305-001', '2024-03-05', 6,  'Tim Teknis Baru (3 Orang) - Onboarding',              3,  'Set tang krimping untuk 3 teknisi baru yang bergabung tim instalasi MyRepublic Solo'),
('2024-03-07 08:00:00', '2024-03-07 08:00:00', 1, 1, 'SO-20240307-001', '2024-03-07', 15, 'Tim Lapangan - Seluruh Tim Instalasi',                40, 'Cable tie untuk merapikan kabel instalasi batch Maret dari berbagai area'),
('2024-03-09 09:00:00', '2024-03-09 09:00:00', 1, 1, 'SO-20240309-001', '2024-03-09', 20, 'Teknisi Lapangan - Klaim Garansi Area Karanganyar',   15, 'Adaptor ONT pengganti untuk pelanggan yang adaptor rusak, wilayah Karanganyar'),
('2024-03-11 10:30:00', '2024-03-11 10:30:00', 1, 1, 'SO-20240311-001', '2024-03-11', 12, 'Proyek LAN - SMK Negeri 2 Surakarta',                 2,  'Kabel UTP Cat6 untuk instalasi jaringan LAN laboratorium komputer SMKN 2 Surakarta'),
('2024-03-13 09:00:00', '2024-03-13 09:00:00', 1, 1, 'SO-20240313-001', '2024-03-13', 13, 'Node Baru - Jl. Ring Road Utara Solo',                2,  'OTB 12 core untuk terminasi kabel FO di node distribusi baru Ring Road Utara'),
('2024-03-15 10:00:00', '2024-03-15 10:00:00', 1, 1, 'SO-20240315-001', '2024-03-15', 16, 'Tim Instalasi Lapangan - Shift Pagi & Siang',         5,  'Helm safety untuk teknisi lapangan, penggantian helm lama yang sudah rusak'),
('2024-03-17 08:30:00', '2024-03-17 08:30:00', 1, 1, 'SO-20240317-001', '2024-03-17', 17, 'Divisi Administrasi & CS - Kantor Solo',              20, 'Kertas HVS untuk cetak dokumen BA instalasi dan laporan bulanan divisi admin'),
('2024-03-19 11:00:00', '2024-03-19 11:00:00', 1, 1, 'SO-20240319-001', '2024-03-19', 18, 'Tim FO - Troubleshooting Kabel Putus Slamet Riyadi',  1,  'Peminjaman OTDR untuk pengukuran dan lokalisir gangguan kabel FO Jl. Slamet Riyadi'),
('2024-03-21 09:30:00', '2024-03-21 09:30:00', 1, 1, 'SO-20240321-001', '2024-03-21', 1,  'Tim Instalasi - Area Gondokusuman & Baturan',         8,  'Router MikroTik untuk instalasi 8 pelanggan baru paket 100Mbps area Gondokusuman');
