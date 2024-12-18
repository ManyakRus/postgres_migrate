--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9
-- Dumped by pg_dump version 14.15 (Ubuntu 14.15-0ubuntu0.22.04.1)

-- Started on 2024-12-18 13:07:59 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 89 (class 2615 OID 676103667)
-- Name: postgres_migrate; Type: SCHEMA; Schema: -; Owner: dev
--

CREATE SCHEMA postgres_migrate;


ALTER SCHEMA postgres_migrate OWNER TO dev;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 461 (class 1259 OID 676124217)
-- Name: postgres_migrate_pg_attribute; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_attribute (
    version_id bigint NOT NULL,
    attrelid oid NOT NULL,
    attname name NOT NULL,
    atttypid oid NOT NULL,
    attstattarget integer NOT NULL,
    attlen smallint NOT NULL,
    attnum smallint NOT NULL,
    attndims integer NOT NULL,
    attcacheoff integer NOT NULL,
    atttypmod integer NOT NULL,
    attbyval boolean NOT NULL,
    attstorage "char" NOT NULL,
    attalign "char" NOT NULL,
    attnotnull boolean NOT NULL,
    atthasdef boolean NOT NULL,
    atthasmissing boolean NOT NULL,
    attidentity "char" NOT NULL,
    attgenerated "char" NOT NULL,
    attisdropped boolean NOT NULL,
    attislocal boolean NOT NULL,
    attinhcount integer NOT NULL,
    attcollation oid NOT NULL
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_attribute OWNER TO dev;

--
-- TOC entry 3707 (class 0 OID 0)
-- Dependencies: 461
-- Name: TABLE postgres_migrate_pg_attribute; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_attribute IS 'В каталоге postgres_migrate_pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в postgres_migrate_pg_attribute существует ровно одна строка.';


--
-- TOC entry 3708 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3709 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attrelid IS 'Таблица, к которой принадлежит столбец';


--
-- TOC entry 3710 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attname; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attname IS 'Имя столбца';


--
-- TOC entry 3711 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.atttypid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.atttypid IS 'Тип данных этого столбца';


--
-- TOC entry 3712 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attstattarget; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attstattarget IS 'Столбец attstattarget управляет детализацией статистики, собираемой по этому столбцу командой ANALYZE. Нулевое значение указывает, что статистика не собирается. При отрицательном значении используется системное ограничение статистики по умолчанию. Точное значение положительных величин определяется типом данных. Для скалярных типов данных, attstattarget задаёт и целевое число собираемых «самых частых значений», и целевое число создаваемых групп гистограммы.';


--
-- TOC entry 3713 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attlen; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attlen IS 'Копия pg_type.typlen из записи типа столбца';


--
-- TOC entry 3714 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attnum; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attnum IS 'Порядковый номер столбца. Обычные столбцы нумеруются по возрастанию, начиная с 1. Системные столбцы, такие как oid, имеют (обычно) отрицательные номера.';


--
-- TOC entry 3715 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attndims; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attndims IS 'Число размерностей, если столбец имеет тип массива; ноль в противном случае. (В настоящее время число размерностей массива не контролируется, поэтому любое ненулевое значение по сути означает «это массив».)';


--
-- TOC entry 3716 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attcacheoff; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attcacheoff IS 'Всегда -1 в постоянном хранилище, но когда запись загружается в память, в этом поле может кешироваться смещение атрибута в строке';


--
-- TOC entry 3717 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.atttypmod; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.atttypmod IS 'В поле atttypmod записывается дополнительное число, связанное с определённым типом данных, указываемое при создании таблицы (например, максимальный размер столбца varchar). Это значение передаётся функциям ввода и преобразования длины конкретного типа. Для типов, которым не нужен atttypmod, это обычно -1.';


--
-- TOC entry 3718 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attbyval; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attbyval IS 'Копия pg_type.typbyval из записи типа столбца';


--
-- TOC entry 3719 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attstorage; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attstorage IS 'Обычно копия pg_type.typstorage из записи типа столбца. Для типов, поддерживающих TOAST, можно изменять это значение после создания столбца и таким образом управлять политикой хранения.';


--
-- TOC entry 3720 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attalign; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attalign IS 'Копия pg_type.typalign из записи типа столбца';


--
-- TOC entry 3721 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attnotnull; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attnotnull IS 'Представляет ограничение NOT NULL.';


--
-- TOC entry 3722 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.atthasdef; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.atthasdef IS 'Столбец имеет значение по умолчанию, в этом случае в каталоге pg_attrdef будет соответствующая запись, определяющая это значение.';


--
-- TOC entry 3723 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attisdropped; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attisdropped IS 'Столбец был удалён и теперь не является рабочим. Удалённый столбец может по-прежнему физически присутствовать в таблице, но анализатор запросов его игнорирует, так что обратиться к нему из SQL нельзя.';


--
-- TOC entry 3724 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attislocal; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attislocal IS 'Столбец определён локально в данном отношении. Заметьте, что столбец может быть определён локально и при этом наследоваться.';


--
-- TOC entry 3725 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attinhcount; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attinhcount IS 'Число прямых предков этого столбца. Столбец с ненулевым числом предков нельзя удалить или переименовать.';


--
-- TOC entry 3726 (class 0 OID 0)
-- Dependencies: 461
-- Name: COLUMN postgres_migrate_pg_attribute.attcollation; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_attribute.attcollation IS 'Заданное для столбца правило сортировки, либо ноль, если тип столбца не сортируемый.';


--
-- TOC entry 462 (class 1259 OID 676125388)
-- Name: postgres_migrate_pg_class; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_class (
    version_id bigint NOT NULL,
    oid oid NOT NULL,
    relname name NOT NULL,
    relnamespace oid NOT NULL,
    reltype oid NOT NULL,
    reloftype oid NOT NULL,
    relowner oid NOT NULL,
    relam oid NOT NULL,
    relfilenode oid NOT NULL,
    reltablespace oid NOT NULL,
    relpages integer NOT NULL,
    reltuples real NOT NULL,
    relallvisible integer NOT NULL,
    reltoastrelid oid NOT NULL,
    relhasindex boolean NOT NULL,
    relisshared boolean NOT NULL,
    relpersistence "char" NOT NULL,
    relkind "char" NOT NULL,
    relnatts smallint NOT NULL,
    relchecks smallint NOT NULL,
    relhasrules boolean NOT NULL,
    relhastriggers boolean NOT NULL,
    relhassubclass boolean NOT NULL,
    relrowsecurity boolean NOT NULL,
    relforcerowsecurity boolean NOT NULL,
    relispopulated boolean NOT NULL,
    relreplident "char" NOT NULL,
    relispartition boolean NOT NULL,
    relrewrite oid NOT NULL,
    relfrozenxid xid NOT NULL,
    relminmxid xid NOT NULL
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_class OWNER TO dev;

--
-- TOC entry 3727 (class 0 OID 0)
-- Dependencies: 462
-- Name: TABLE postgres_migrate_pg_class; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_class IS 'В каталоге postgres_migrate_pg_class описываются таблицы и практически всё, что имеет столбцы или каким-то образом подобно таблице. Сюда входят индексы (но смотрите также postgres_migrate_pg_index), последовательности, представления, материализованные представления, составные типы и таблицы TOAST; см. relkind. Далее, подразумевая все эти типы объектов, мы будем говорить об «отношениях». Не все столбцы здесь имеют смысл для всех типов отношений.';


--
-- TOC entry 3728 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3729 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.oid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.oid IS 'Идентификатор строки (скрытый атрибут; должен выбираться явно)';


--
-- TOC entry 3730 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relname; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relname IS 'Имя таблицы, индекса, представления и т. п.';


--
-- TOC entry 3731 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relnamespace; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relnamespace IS 'OID пространства имён, содержащего это отношение';


--
-- TOC entry 3732 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.reltype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.reltype IS 'OID типа данных, соответствующего типу строки этой таблицы, если таковой есть (ноль для индексов, так как они не имеют записи в pg_type)';


--
-- TOC entry 3733 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.reloftype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.reloftype IS 'Для типизированных таблиц, OID нижележащего составного типа, или ноль для всех других отношений';


--
-- TOC entry 3734 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relowner; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relowner IS 'Владелец отношения';


--
-- TOC entry 3735 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relam; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relam IS 'Если это индекс, применяемый метод доступа (B-дерево, хеш и т. д.)';


--
-- TOC entry 3736 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relfilenode; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relfilenode IS 'Имя файла на диске с этим отношением; ноль означает, что это «отображённое» представление, имя файла для которого определяется состоянием на нижнем уровне';


--
-- TOC entry 3737 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.reltablespace; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.reltablespace IS 'Табличное пространство, в котором хранится это отношение. Если ноль, подразумевается пространство базы данных по умолчанию. (Не имеет значения, если с отношением не связан файл на диске.)';


--
-- TOC entry 3738 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relpages; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relpages IS 'Размер представления этой таблицы на диске (в страницах размера BLCKSZ). Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.';


--
-- TOC entry 3739 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.reltuples; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.reltuples IS 'Число строк в таблице. Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.';


--
-- TOC entry 3740 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relallvisible; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relallvisible IS 'Число страниц, помеченных как «полностью видимые» в карте видимости таблицы. Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.';


--
-- TOC entry 3741 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.reltoastrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.reltoastrelid IS 'OID таблицы TOAST, связанной с данной таблицей, или 0, если таковой нет. В таблицу TOAST, как во вторичную, «выносятся» большие атрибуты.';


--
-- TOC entry 3742 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relhasindex; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relhasindex IS 'True, если это таблица и она имеет (или недавно имела) индексы';


--
-- TOC entry 3743 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relisshared; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relisshared IS 'True, если эта таблица разделяется всеми базами данных в кластере. Разделяемыми являются только некоторые системные каталоги (как например, pg_database).';


--
-- TOC entry 3744 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relpersistence; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relpersistence IS 'p = постоянная таблица (permanent), u = нежурналируемая таблица (unlogged), t = временная таблица (temporary)';


--
-- TOC entry 3745 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relkind; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relkind IS 'r = обычная таблица, i = индекс (index), S = последовательность (sequence), v = представление (view), m = материализованное представление (materialized view), c = составной тип (composite), t = таблица TOAST, f = сторонняя таблица (foreign)';


--
-- TOC entry 3746 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relnatts; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relnatts IS 'Число пользовательских столбцов в отношении (системные столбцы не считаются). Столько же соответствующих строк должно быть в postgres_migrate_pg_attribute. См. также postgres_migrate_pg_attribute.attnum.';


--
-- TOC entry 3747 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relchecks; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relchecks IS 'Число ограничений CHECK в таблице; см. каталог postgres_migrate_pg_constraint';


--
-- TOC entry 3748 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relhasrules; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relhasrules IS 'True, если для таблицы определены (или были определены) правила; см. каталог pg_rewrite';


--
-- TOC entry 3749 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relhastriggers; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relhastriggers IS 'True, если для таблицы определены (или были определены) триггеры; см. каталог pg_trigger';


--
-- TOC entry 3750 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relhassubclass; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relhassubclass IS 'True, если у таблицы есть (или были) потомки в иерархии наследования';


--
-- TOC entry 3751 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relrowsecurity; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relrowsecurity IS 'True, если для таблицы включена защита на уровне строк; см. каталог pg_policy';


--
-- TOC entry 3752 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relforcerowsecurity; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relforcerowsecurity IS 'True, если защита на уровне строк (когда она включена) также применяется к владельцу таблицы; см. каталог pg_policy';


--
-- TOC entry 3753 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relispopulated; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relispopulated IS 'True, если отношение наполнено данными (это истинно для всех отношений, кроме некоторых материализованных представлений)';


--
-- TOC entry 3754 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relreplident; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relreplident IS 'Столбцы, формирующие «идентификатор реплики» для строк: d = по умолчанию (первичный ключ, если есть), n = никакие (nothing), f = все столбцы, i = индекс со свойством indisreplident (если ранее использованный индекс удалён, действует так же, как n)';


--
-- TOC entry 3755 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relfrozenxid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relfrozenxid IS 'Идентификаторы транзакций, предшествующие данному, в этой таблице заменены постоянным («замороженным») идентификатором транзакции. Это нужно для определения, когда требуется очищать таблицу для предотвращения зацикливания идентификаторов или для сокращения объёма pg_clog. Если это отношение — не таблица, значение равно нулю (InvalidTransactionId).';


--
-- TOC entry 3756 (class 0 OID 0)
-- Dependencies: 462
-- Name: COLUMN postgres_migrate_pg_class.relminmxid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_class.relminmxid IS 'Идентификаторы мультитранзакций, предшествующие данному, в этой таблице заменены другим идентификатором транзакции. Это нужно для определения, когда требуется очищать таблицу для предотвращения зацикливания идентификаторов мультитранзакций или для сокращения объёма pg_multixact. Если это отношение — не таблица, значение равно нулю (InvalidMultiXactId).';


--
-- TOC entry 463 (class 1259 OID 676125405)
-- Name: postgres_migrate_pg_constraint; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_constraint (
    version_id bigint NOT NULL,
    oid oid NOT NULL,
    conname name NOT NULL,
    connamespace oid NOT NULL,
    contype "char" NOT NULL,
    condeferrable boolean NOT NULL,
    condeferred boolean NOT NULL,
    convalidated boolean NOT NULL,
    conrelid oid NOT NULL,
    contypid oid NOT NULL,
    conindid oid NOT NULL,
    conparentid oid NOT NULL,
    confrelid oid NOT NULL,
    confupdtype "char" NOT NULL,
    confdeltype "char" NOT NULL,
    confmatchtype "char" NOT NULL,
    conislocal boolean NOT NULL,
    coninhcount integer NOT NULL,
    connoinherit boolean NOT NULL,
    conkey smallint[],
    confkey smallint[],
    conpfeqop oid[],
    conppeqop oid[],
    conffeqop oid[],
    conexclop oid[]
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_constraint OWNER TO dev;

--
-- TOC entry 3757 (class 0 OID 0)
-- Dependencies: 463
-- Name: TABLE postgres_migrate_pg_constraint; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_constraint IS 'В каталоге postgres_migrate_pg_constraint хранятся ограничения-проверки, ограничения-исключения, а также ограничения первичного ключа, уникальности и внешних ключей, определённые для таблиц. (Ограничения столбцов описываются как и все остальные. Любое ограничение столбца равнозначно некоторому ограничению таблицы.) Ограничения на NULL представляются не здесь, а в каталоге postgres_migrate_pg_attribute.

Для пользовательских триггеров ограничений (создаваемых командой CREATE CONSTRAINT TRIGGER) в этой таблице также создаётся запись.

Здесь также хранятся ограничения доменов.';


--
-- TOC entry 3758 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3759 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.oid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.oid IS 'Идентификатор строки (скрытый атрибут; должен выбираться явно)';


--
-- TOC entry 3760 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conname; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conname IS 'Имя ограничения (не обязательно уникальное!)';


--
-- TOC entry 3761 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.connamespace; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.connamespace IS 'OID пространства имён, содержащего это ограничение';


--
-- TOC entry 3762 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.contype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.contype IS 'c = ограничение-проверка (check), f = внешний ключ (foreign key), p = первичный ключ (primary key), u = ограничение уникальности (unique), t = триггер ограничения (trigger), x = ограничение-исключение (exclusion)';


--
-- TOC entry 3763 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.condeferrable; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.condeferrable IS 'Является ли ограничение откладываемым?';


--
-- TOC entry 3764 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.condeferred; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.condeferred IS 'Является ли ограничение отложенным по умолчанию?';


--
-- TOC entry 3765 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.convalidated; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.convalidated IS 'Было ли ограничение проверено? В настоящее время значение false возможно только для внешних ключей и ограничений CHECK';


--
-- TOC entry 3766 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conrelid IS 'Таблица, для которой установлено это ограничение; 0, если это не ограничение таблицы';


--
-- TOC entry 3767 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.contypid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.contypid IS 'Домен, к которому относится это ограничение; 0, если это не ограничение домена';


--
-- TOC entry 3768 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conindid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conindid IS 'Индекс, поддерживающий это ограничение, если это ограничение уникальности, первичного или внешнего ключа, либо ограничение-исключение; в противном случае — 0';


--
-- TOC entry 3769 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conparentid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conparentid IS 'Соответствующее ограничение в родительской секционированной таблице, если это ограничение в секции; иначе ноль';


--
-- TOC entry 3770 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.confrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.confrelid IS 'Если это внешний ключ, таблица, на которую он ссылается; иначе 0';


--
-- TOC entry 3771 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.confupdtype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.confupdtype IS 'Код действия при изменении внешнего ключа: a = нет действия, r = ограничить (restrict), c = каскадное действие (cascade), n = присвоить NULL, d = поведение по умолчанию';


--
-- TOC entry 3772 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.confdeltype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.confdeltype IS 'Код действия при удалении внешнего ключа: a = нет действия, r = ограничить (restrict), c = каскадное действие (cascade), n = присвоить NULL, d = поведение по умолчанию';


--
-- TOC entry 3773 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.confmatchtype; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.confmatchtype IS 'Тип сопоставления внешнего ключа: f = полное (full), p = частичное (partial), s = простое (simple)';


--
-- TOC entry 3774 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conislocal; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conislocal IS 'Ограничение определено локально в данном отношении. Заметьте, что ограничение может быть определено локально и при этом наследоваться.';


--
-- TOC entry 3775 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.coninhcount; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.coninhcount IS 'Число прямых предков этого ограничения. Ограничение с ненулевым числом предков нельзя удалить или переименовать.';


--
-- TOC entry 3776 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.connoinherit; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.connoinherit IS 'Ограничение определено локально для данного отношения и является ненаследуемым.';


--
-- TOC entry 3777 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conkey; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conkey IS 'Для ограничений таблицы (включая внешние ключи, но не триггеры ограничений), определяет список столбцов, образующих ограничение';


--
-- TOC entry 3778 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.confkey; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.confkey IS 'Для внешнего ключа определяет список столбцов, на которые он ссылается';


--
-- TOC entry 3779 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conpfeqop; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conpfeqop IS 'Для внешнего ключа — список операторов равенства для сравнений PK = FK';


--
-- TOC entry 3780 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conppeqop; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conppeqop IS 'Для внешнего ключа — список операторов равенства для сравнений PK = PK';


--
-- TOC entry 3781 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conffeqop; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conffeqop IS 'Для внешнего ключа — список операторов равенства для сравнений FK = FK';


--
-- TOC entry 3782 (class 0 OID 0)
-- Dependencies: 463
-- Name: COLUMN postgres_migrate_pg_constraint.conexclop; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_constraint.conexclop IS 'Для ограничения-исключения — список операторов исключения по столбцам';


--
-- TOC entry 464 (class 1259 OID 676125433)
-- Name: postgres_migrate_pg_description; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_description (
    version_id bigint NOT NULL,
    objoid oid NOT NULL,
    classoid oid NOT NULL,
    objsubid integer NOT NULL,
    description text NOT NULL COLLATE pg_catalog."C"
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_description OWNER TO dev;

--
-- TOC entry 3783 (class 0 OID 0)
-- Dependencies: 464
-- Name: TABLE postgres_migrate_pg_description; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_description IS 'В каталоге postgres_migrate_pg_description хранятся дополнительные описания (комментарии) для объектов баз данных. Описания можно задавать с помощью команды COMMENT и просматривать в psql, используя команды \d. В начальном содержимом postgres_migrate_pg_description представлены описания многих встроенных системных объектов.

Также смотрите каталог pg_shdescription, который играет подобную роль в отношении совместно используемых объектов в кластере баз данных.';


--
-- TOC entry 3784 (class 0 OID 0)
-- Dependencies: 464
-- Name: COLUMN postgres_migrate_pg_description.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_description.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3785 (class 0 OID 0)
-- Dependencies: 464
-- Name: COLUMN postgres_migrate_pg_description.objoid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_description.objoid IS 'OID объекта, к которому относится это описание';


--
-- TOC entry 3786 (class 0 OID 0)
-- Dependencies: 464
-- Name: COLUMN postgres_migrate_pg_description.classoid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_description.classoid IS 'OID системного каталога, к которому относится этот объект';


--
-- TOC entry 3787 (class 0 OID 0)
-- Dependencies: 464
-- Name: COLUMN postgres_migrate_pg_description.objsubid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_description.objsubid IS 'Для комментария к столбцу таблицы это номер столбца (objoid и classoid указывают на саму таблицу). Для всех других типов объектов это поле содержит ноль.';


--
-- TOC entry 3788 (class 0 OID 0)
-- Dependencies: 464
-- Name: COLUMN postgres_migrate_pg_description.description; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_description.description IS 'Произвольный текст, служащий описанием данного объекта';


--
-- TOC entry 465 (class 1259 OID 676125440)
-- Name: postgres_migrate_pg_index; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_index (
    version_id bigint NOT NULL,
    indexrelid oid NOT NULL,
    indrelid oid NOT NULL,
    indnatts smallint NOT NULL,
    indnkeyatts smallint NOT NULL,
    indisunique boolean NOT NULL,
    indisprimary boolean NOT NULL,
    indisexclusion boolean NOT NULL,
    indimmediate boolean NOT NULL,
    indisclustered boolean NOT NULL,
    indisvalid boolean NOT NULL,
    indcheckxmin boolean NOT NULL,
    indisready boolean NOT NULL,
    indislive boolean NOT NULL,
    indisreplident boolean NOT NULL,
    indkey int2vector NOT NULL,
    indcollation oidvector NOT NULL,
    indclass oidvector NOT NULL,
    indoption int2vector NOT NULL,
    indexprs pg_node_tree COLLATE pg_catalog."C",
    indpred pg_node_tree COLLATE pg_catalog."C"
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_index OWNER TO dev;

--
-- TOC entry 3789 (class 0 OID 0)
-- Dependencies: 465
-- Name: TABLE postgres_migrate_pg_index; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_index IS 'В каталоге postgres_migrate_pg_index содержится часть информации об индексах. Остальная информация в основном находится в postgres_migrate_pg_class.';


--
-- TOC entry 3790 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3791 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indexrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indexrelid IS 'OID записи в postgres_migrate_pg_class для этого индекса';


--
-- TOC entry 3792 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indrelid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indrelid IS 'OID записи в postgres_migrate_pg_class для таблицы, к которой относится этот индекс';


--
-- TOC entry 3793 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indnatts; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indnatts IS 'Число столбцов в индексе (повторяет значение postgres_migrate_pg_class.relnatts)';


--
-- TOC entry 3794 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indnkeyatts; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indnkeyatts IS 'Количество ключевых столбцов в индексе. «Ключевые столбцы» это обычные столбцы, в отличие от «включённых» столбцов.';


--
-- TOC entry 3795 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisunique; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisunique IS 'Если true, это уникальный индекс';


--
-- TOC entry 3796 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisprimary; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisprimary IS 'Если true, этот индекс представляет первичный ключ таблицы (в этом случае и в поле indisunique должно быть значение true)';


--
-- TOC entry 3797 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisexclusion; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisexclusion IS 'Если true, этот индекс поддерживает ограничение-исключение';


--
-- TOC entry 3798 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indimmediate; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indimmediate IS 'Если true, проверка уникальности осуществляется непосредственно при добавлении данных (неприменимо, если значение indisunique не true)';


--
-- TOC entry 3799 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisclustered; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisclustered IS 'Если true, таблица в последний раз кластеризовалась по этому индексу';


--
-- TOC entry 3800 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisvalid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisvalid IS 'Если true, индекс можно применять в запросах. Значение false означает, что индекс, возможно, неполный: он будет тем не менее изменяться командами INSERT/UPDATE, но безопасно применять его в запросах нельзя. Если он уникальный, свойство уникальности также не гарантируется.';


--
-- TOC entry 3801 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indcheckxmin; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indcheckxmin IS 'Если true, запросы не должны использовать этот индекс, пока поле xmin данной записи в postgres_migrate_pg_index не окажется ниже их горизонта событий TransactionXmin, так как таблица может содержать оборванные цепочки HOT, с видимыми несовместимыми строками';


--
-- TOC entry 3802 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisready; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisready IS 'Если true, индекс готов к добавлению данных. Значение false означает, что индекс игнорируется операциями INSERT/UPDATE.';


--
-- TOC entry 3803 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indislive; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indislive IS 'Если false, индекс находится в процессе удаления и его следует игнорировать для любых целей (включая вопрос применимости HOT)';


--
-- TOC entry 3804 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indisreplident; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indisreplident IS 'Если true, этот индекс выбран в качестве «идентификатора реплики» командой ALTER TABLE ... REPLICA IDENTITY USING INDEX ...';


--
-- TOC entry 3805 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indkey; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indkey IS 'Это массив из indnatts значений, указывающих, какие столбцы таблицы индексирует этот индекс. Например, значения 1 3 будут означать, что ключ индекса составляют первый и третий столбцы таблицы. Ноль в этом массиве означает, что соответствующий атрибут индекса определяется выражением со столбцами таблицы, а не просто ссылкой на столбец.';


--
-- TOC entry 3806 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indcollation; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indcollation IS 'Для каждого столбца в ключе индекса этот массив содержит OID правила сортировки для применения в этом индексе.';


--
-- TOC entry 3807 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indclass; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indclass IS 'Для каждого столбца в ключе индекса этот массив содержит OID применяемых классов операторов. Подробнее это рассматривается в описании pg_opclass.';


--
-- TOC entry 3808 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indoption; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indoption IS 'Это массив из indnatts значений, в которых хранятся битовые флаги для отдельных столбцов. Значение этих флагов определяется методом доступа конкретного индекса.';


--
-- TOC entry 3809 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indexprs; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indexprs IS 'Деревья выражений (в представлении nodeToString()) для атрибутов индекса, не являющихся простыми ссылками на столбцы. Этот список содержит один элемент для каждого нулевого значения в indkey. Значением может быть NULL, если все атрибуты индекса представляют собой простые ссылки.';


--
-- TOC entry 3810 (class 0 OID 0)
-- Dependencies: 465
-- Name: COLUMN postgres_migrate_pg_index.indpred; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_index.indpred IS 'Дерево выражения (в представлении nodeToString()) для предиката частичного индекса, либо NULL, если это не частичный индекс.';


--
-- TOC entry 466 (class 1259 OID 676125454)
-- Name: postgres_migrate_pg_namespace; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_pg_namespace (
    version_id bigint NOT NULL,
    oid oid NOT NULL,
    nspname name NOT NULL,
    nspowner oid NOT NULL,
    nspacl aclitem[]
);


ALTER TABLE postgres_migrate.postgres_migrate_pg_namespace OWNER TO dev;

--
-- TOC entry 3811 (class 0 OID 0)
-- Dependencies: 466
-- Name: TABLE postgres_migrate_pg_namespace; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON TABLE postgres_migrate.postgres_migrate_pg_namespace IS 'В postgres_migrate_pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.';


--
-- TOC entry 3812 (class 0 OID 0)
-- Dependencies: 466
-- Name: COLUMN postgres_migrate_pg_namespace.version_id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_namespace.version_id IS 'Версия изменений (ИД)';


--
-- TOC entry 3813 (class 0 OID 0)
-- Dependencies: 466
-- Name: COLUMN postgres_migrate_pg_namespace.oid; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_namespace.oid IS 'Идентификатор строки (скрытый атрибут; должен выбираться явно)';


--
-- TOC entry 3814 (class 0 OID 0)
-- Dependencies: 466
-- Name: COLUMN postgres_migrate_pg_namespace.nspname; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_namespace.nspname IS 'Имя пространства имён';


--
-- TOC entry 3815 (class 0 OID 0)
-- Dependencies: 466
-- Name: COLUMN postgres_migrate_pg_namespace.nspowner; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_namespace.nspowner IS 'Владелец пространства имён';


--
-- TOC entry 3816 (class 0 OID 0)
-- Dependencies: 466
-- Name: COLUMN postgres_migrate_pg_namespace.nspacl; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_pg_namespace.nspacl IS 'Права доступа; за подробностями обратитесь к описанию GRANT и REVOKE';


--
-- TOC entry 460 (class 1259 OID 676104923)
-- Name: version; Type: TABLE; Schema: postgres_migrate; Owner: dev
--

CREATE TABLE postgres_migrate.postgres_migrate_version (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    modified_at timestamp with time zone,
    name text,
    description text
);


ALTER TABLE postgres_migrate.postgres_migrate_version OWNER TO dev;

--
-- TOC entry 3817 (class 0 OID 0)
-- Dependencies: 460
-- Name: COLUMN version.id; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_version.id IS 'Уникальный технический идентификатор';


--
-- TOC entry 3818 (class 0 OID 0)
-- Dependencies: 460
-- Name: COLUMN version.created_at; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_version.created_at IS 'Дата создания элемента';


--
-- TOC entry 3819 (class 0 OID 0)
-- Dependencies: 460
-- Name: COLUMN version.modified_at; Type: COMMENT; Schema: postgres_migrate; Owner: dev
--

COMMENT ON COLUMN postgres_migrate.postgres_migrate_version.modified_at IS 'Дата последнего изменения элемента';


--
-- TOC entry 459 (class 1259 OID 676104921)
-- Name: version_id_seq; Type: SEQUENCE; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE postgres_migrate.postgres_migrate_version ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME postgres_migrate.postgres_migrate_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3543 (class 2606 OID 676182095)
-- Name: postgres_migrate_pg_attribute postgres_migrate_pg_attribute_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_attribute
    ADD CONSTRAINT postgres_migrate_pg_attribute_pk PRIMARY KEY (attrelid, attname, version_id);


--
-- TOC entry 3548 (class 2606 OID 676184153)
-- Name: postgres_migrate_pg_class postgres_migrate_pg_class_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_class
    ADD CONSTRAINT postgres_migrate_pg_class_pk PRIMARY KEY (oid, version_id);


--
-- TOC entry 3557 (class 2606 OID 676185121)
-- Name: postgres_migrate_pg_constraint postgres_migrate_pg_constraint_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_constraint
    ADD CONSTRAINT postgres_migrate_pg_constraint_pk PRIMARY KEY (oid, version_id);


--
-- TOC entry 3560 (class 2606 OID 676167159)
-- Name: postgres_migrate_pg_description postgres_migrate_pg_description_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_description
    ADD CONSTRAINT postgres_migrate_pg_description_pk PRIMARY KEY (objoid, classoid, objsubid, version_id);


--
-- TOC entry 3564 (class 2606 OID 676187628)
-- Name: postgres_migrate_pg_index postgres_migrate_pg_index_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_index
    ADD CONSTRAINT postgres_migrate_pg_index_pk PRIMARY KEY (indrelid, indexrelid, version_id);


--
-- TOC entry 3566 (class 2606 OID 676167165)
-- Name: postgres_migrate_pg_namespace pg_name_space_pk; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_namespace
    ADD CONSTRAINT pg_name_space_pk PRIMARY KEY (oid, version_id);


--
-- TOC entry 3541 (class 2606 OID 676104930)
-- Name: version version_pkey; Type: CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_version
    ADD CONSTRAINT version_pkey PRIMARY KEY (id);


--
-- TOC entry 3544 (class 1259 OID 676125335)
-- Name: postgres_migrate_pg_attribute_relid_attnam_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_attribute_relid_attnam_index ON postgres_migrate.postgres_migrate_pg_attribute USING btree (attrelid, attname);


--
-- TOC entry 3545 (class 1259 OID 676125372)
-- Name: postgres_migrate_pg_attribute_relid_attnum_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_attribute_relid_attnum_index ON postgres_migrate.postgres_migrate_pg_attribute USING btree (attrelid, attnum);


--
-- TOC entry 3546 (class 1259 OID 676125399)
-- Name: postgres_migrate_pg_class_oid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_class_oid_index ON postgres_migrate.postgres_migrate_pg_class USING btree (oid);


--
-- TOC entry 3549 (class 1259 OID 676125400)
-- Name: postgres_migrate_pg_class_relname_nsp_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_class_relname_nsp_index ON postgres_migrate.postgres_migrate_pg_class USING btree (relname, relnamespace);


--
-- TOC entry 3550 (class 1259 OID 676125401)
-- Name: postgres_migrate_pg_class_tblspc_relfilenode_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE INDEX postgres_migrate_pg_class_tblspc_relfilenode_index ON postgres_migrate.postgres_migrate_pg_class USING btree (reltablespace, relfilenode);


--
-- TOC entry 3551 (class 1259 OID 676125416)
-- Name: postgres_migrate_pg_constraint_conname_nsp_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE INDEX postgres_migrate_pg_constraint_conname_nsp_index ON postgres_migrate.postgres_migrate_pg_constraint USING btree (conname, connamespace);


--
-- TOC entry 3552 (class 1259 OID 676125417)
-- Name: postgres_migrate_pg_constraint_conparentid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE INDEX postgres_migrate_pg_constraint_conparentid_index ON postgres_migrate.postgres_migrate_pg_constraint USING btree (conparentid);


--
-- TOC entry 3553 (class 1259 OID 676125418)
-- Name: postgres_migrate_pg_constraint_conrelid_contypid_conname_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_constraint_conrelid_contypid_conname_index ON postgres_migrate.postgres_migrate_pg_constraint USING btree (conrelid, contypid, conname);


--
-- TOC entry 3554 (class 1259 OID 676125419)
-- Name: postgres_migrate_pg_constraint_contypid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE INDEX postgres_migrate_pg_constraint_contypid_index ON postgres_migrate.postgres_migrate_pg_constraint USING btree (contypid);


--
-- TOC entry 3555 (class 1259 OID 676125420)
-- Name: postgres_migrate_pg_constraint_oid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_constraint_oid_index ON postgres_migrate.postgres_migrate_pg_constraint USING btree (oid);


--
-- TOC entry 3558 (class 1259 OID 676125439)
-- Name: postgres_migrate_pg_description_o_c_o_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_description_o_c_o_index ON postgres_migrate.postgres_migrate_pg_description USING btree (objoid, classoid, objsubid);


--
-- TOC entry 3561 (class 1259 OID 676125451)
-- Name: postgres_migrate_pg_index_indexrelid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_index_indexrelid_index ON postgres_migrate.postgres_migrate_pg_index USING btree (indexrelid);


--
-- TOC entry 3562 (class 1259 OID 676125452)
-- Name: postgres_migrate_pg_index_indrelid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE INDEX postgres_migrate_pg_index_indrelid_index ON postgres_migrate.postgres_migrate_pg_index USING btree (indrelid);


--
-- TOC entry 3567 (class 1259 OID 676125465)
-- Name: postgres_migrate_pg_namespace_nspname_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_namespace_nspname_index ON postgres_migrate.postgres_migrate_pg_namespace USING btree (nspname);


--
-- TOC entry 3568 (class 1259 OID 676125466)
-- Name: postgres_migrate_pg_namespace_oid_index; Type: INDEX; Schema: postgres_migrate; Owner: dev
--

CREATE UNIQUE INDEX postgres_migrate_pg_namespace_oid_index ON postgres_migrate.postgres_migrate_pg_namespace USING btree (oid);


--
-- TOC entry 3569 (class 2606 OID 676124277)
-- Name: postgres_migrate_pg_attribute postgres_migrate_pg_attribute_version_id_fk; Type: FK CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_attribute
    ADD CONSTRAINT postgres_migrate_pg_attribute_version_id_fk FOREIGN KEY (version_id) REFERENCES postgres_migrate.postgres_migrate_version(id);


--
-- TOC entry 3570 (class 2606 OID 676125394)
-- Name: postgres_migrate_pg_class postgres_migrate_pg_class_version_id_fk; Type: FK CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_class
    ADD CONSTRAINT postgres_migrate_pg_class_version_id_fk FOREIGN KEY (version_id) REFERENCES postgres_migrate.postgres_migrate_version(id);


--
-- TOC entry 3571 (class 2606 OID 676125411)
-- Name: postgres_migrate_pg_constraint postgres_migrate_pg_constraint_version_id_fk; Type: FK CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_constraint
    ADD CONSTRAINT postgres_migrate_pg_constraint_version_id_fk FOREIGN KEY (version_id) REFERENCES postgres_migrate.postgres_migrate_version(id);


--
-- TOC entry 3572 (class 2606 OID 676125446)
-- Name: postgres_migrate_pg_index postgres_migrate_pg_index_version_id_fk; Type: FK CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_index
    ADD CONSTRAINT postgres_migrate_pg_index_version_id_fk FOREIGN KEY (version_id) REFERENCES postgres_migrate.postgres_migrate_version(id);


--
-- TOC entry 3573 (class 2606 OID 676125460)
-- Name: postgres_migrate_pg_namespace postgres_migrate_pg_namespace_version_id_fk; Type: FK CONSTRAINT; Schema: postgres_migrate; Owner: dev
--

ALTER TABLE ONLY postgres_migrate.postgres_migrate_pg_namespace
    ADD CONSTRAINT postgres_migrate_pg_namespace_version_id_fk FOREIGN KEY (version_id) REFERENCES postgres_migrate.postgres_migrate_version(id);


-- Completed on 2024-12-18 13:07:59 MSK

--
-- PostgreSQL database dump complete
--

