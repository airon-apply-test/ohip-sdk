/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ArticlePostItType } from './ArticlePostItType';
import type { CurrencyAmountType } from './CurrencyAmountType';
/**
 * Information regarding an article.
 * @export
 * @interface ArticleInfoType
 */
export interface ArticleInfoType {
    /**
     * Description of the article.
     * @type {string}
     * @memberof ArticleInfoType
     */
    description?: string;
    /**
     * Transaction code to which the article belongs.
     * @type {string}
     * @memberof ArticleInfoType
     */
    transactionCode?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof ArticleInfoType
     */
    price?: CurrencyAmountType;
    /**
     *
     * @type {ArticlePostItType}
     * @memberof ArticleInfoType
     */
    postIt?: ArticlePostItType;
    /**
     * Unique Universal product code of the article.
     * @type {string}
     * @memberof ArticleInfoType
     */
    universalProductCode?: string;
    /**
     * Hotel code to which the article belongs.
     * @type {string}
     * @memberof ArticleInfoType
     */
    hotelId?: string;
    /**
     * Unique code of the article.
     * @type {string}
     * @memberof ArticleInfoType
     */
    articleCode?: string;
    /**
     * Indicates whether the article is inactive or not.
     * @type {boolean}
     * @memberof ArticleInfoType
     */
    inactive?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ArticleInfoType
     */
    orderSequence?: number;
}
/**
 * Check if a given object implements the ArticleInfoType interface.
 */
export declare function instanceOfArticleInfoType(value: object): boolean;
export declare function ArticleInfoTypeFromJSON(json: any): ArticleInfoType;
export declare function ArticleInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ArticleInfoType;
export declare function ArticleInfoTypeToJSON(value?: ArticleInfoType | null): any;
