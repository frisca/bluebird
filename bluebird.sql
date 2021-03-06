PGDMP         0                x            bluebird    11.5    11.5     P           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            Q           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            R           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            S           1262    17792    bluebird    DATABASE     f   CREATE DATABASE bluebird WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C' LC_CTYPE = 'C';
    DROP DATABASE bluebird;
             postgres    false            �            1259    17808    orders    TABLE     �   CREATE TABLE public.orders (
    id bigint NOT NULL,
    order_no character varying(100),
    total integer,
    created_date timestamp without time zone,
    created_by character varying(50),
    user_id integer,
    reff_no character varying(50)
);
    DROP TABLE public.orders;
       public         postgres    false            �            1259    17816    orders_detail    TABLE       CREATE TABLE public.orders_detail (
    id bigint NOT NULL,
    order_id character varying(100),
    product character varying(100),
    price integer,
    qty integer,
    created_date timestamp without time zone,
    created_by character varying(50),
    user_id integer
);
 !   DROP TABLE public.orders_detail;
       public         postgres    false            �            1259    17814    orders_detail_id_seq    SEQUENCE     }   CREATE SEQUENCE public.orders_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.orders_detail_id_seq;
       public       postgres    false    201            T           0    0    orders_detail_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.orders_detail_id_seq OWNED BY public.orders_detail.id;
            public       postgres    false    200            �            1259    17806    orders_id_seq    SEQUENCE     v   CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public       postgres    false    199            U           0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
            public       postgres    false    198            �            1259    17795    user    TABLE        CREATE TABLE public."user" (
    id bigint NOT NULL,
    username character varying(50),
    password character varying(50)
);
    DROP TABLE public."user";
       public         postgres    false            �            1259    17793    user_id_seq    SEQUENCE     t   CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.user_id_seq;
       public       postgres    false    197            V           0    0    user_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;
            public       postgres    false    196            �           2604    17811 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    198    199    199            �           2604    17819    orders_detail id    DEFAULT     t   ALTER TABLE ONLY public.orders_detail ALTER COLUMN id SET DEFAULT nextval('public.orders_detail_id_seq'::regclass);
 ?   ALTER TABLE public.orders_detail ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    201    200    201            �           2604    17798    user id    DEFAULT     d   ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);
 8   ALTER TABLE public."user" ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    196    197    197            K          0    17808    orders 
   TABLE DATA               a   COPY public.orders (id, order_no, total, created_date, created_by, user_id, reff_no) FROM stdin;
    public       postgres    false    199   �       M          0    17816    orders_detail 
   TABLE DATA               m   COPY public.orders_detail (id, order_id, product, price, qty, created_date, created_by, user_id) FROM stdin;
    public       postgres    false    201   �       I          0    17795    user 
   TABLE DATA               8   COPY public."user" (id, username, password) FROM stdin;
    public       postgres    false    197   :       W           0    0    orders_detail_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.orders_detail_id_seq', 4, true);
            public       postgres    false    200            X           0    0    orders_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.orders_id_seq', 12, true);
            public       postgres    false    198            Y           0    0    user_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.user_id_seq', 1, true);
            public       postgres    false    196            �           2606    17821     orders_detail orders_detail_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.orders_detail
    ADD CONSTRAINT orders_detail_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.orders_detail DROP CONSTRAINT orders_detail_pkey;
       public         postgres    false    201            �           2606    17813    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public         postgres    false    199            �           2606    17800    user user_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public."user" DROP CONSTRAINT user_pkey;
       public         postgres    false    197            K   7   x�34���3000�4�� �.)X�'g�X��̌�����H1z\\\ �ka      M   E   x�3����K���LUp�/J�K�4500�4�� .#NCBj�9	�1�44��	JM*-�P���� �7$�      I      x�3�L�)MM�,J�442615����� M�t     