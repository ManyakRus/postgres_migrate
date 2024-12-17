-- Table: pg_catalog.pg_constraint

-- DROP TABLE IF EXISTS pg_catalog.pg_constraint;

CREATE TABLE IF NOT EXISTS pg_catalog.pg_constraint
(
    oid oid NOT NULL,
    conname name COLLATE pg_catalog."C" NOT NULL,
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
    conexclop oid[],
    conbin pg_node_tree COLLATE pg_catalog."C"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS pg_catalog.pg_constraint
    OWNER to postgres;

REVOKE ALL ON TABLE pg_catalog.pg_constraint FROM PUBLIC;

GRANT SELECT ON TABLE pg_catalog.pg_constraint TO PUBLIC;

GRANT ALL ON TABLE pg_catalog.pg_constraint TO postgres;
-- Index: pg_constraint_conname_nsp_index

-- DROP INDEX IF EXISTS pg_catalog.pg_constraint_conname_nsp_index;

CREATE INDEX IF NOT EXISTS pg_constraint_conname_nsp_index
    ON pg_catalog.pg_constraint USING btree
    (conname COLLATE pg_catalog."C" ASC NULLS LAST, connamespace ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: pg_constraint_conparentid_index

-- DROP INDEX IF EXISTS pg_catalog.pg_constraint_conparentid_index;

CREATE INDEX IF NOT EXISTS pg_constraint_conparentid_index
    ON pg_catalog.pg_constraint USING btree
    (conparentid ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: pg_constraint_conrelid_contypid_conname_index

-- DROP INDEX IF EXISTS pg_catalog.pg_constraint_conrelid_contypid_conname_index;

CREATE UNIQUE INDEX IF NOT EXISTS pg_constraint_conrelid_contypid_conname_index
    ON pg_catalog.pg_constraint USING btree
    (conrelid ASC NULLS LAST, contypid ASC NULLS LAST, conname COLLATE pg_catalog."C" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: pg_constraint_contypid_index

-- DROP INDEX IF EXISTS pg_catalog.pg_constraint_contypid_index;

CREATE INDEX IF NOT EXISTS pg_constraint_contypid_index
    ON pg_catalog.pg_constraint USING btree
    (contypid ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: pg_constraint_oid_index

-- DROP INDEX IF EXISTS pg_catalog.pg_constraint_oid_index;

CREATE UNIQUE INDEX IF NOT EXISTS pg_constraint_oid_index
    ON pg_catalog.pg_constraint USING btree
    (oid ASC NULLS LAST)
    TABLESPACE pg_default;