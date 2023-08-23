/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ArticleInfoType } from './ArticleInfoType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to change articles.
 * @export
 * @interface Articles
 */
export interface Articles {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof Articles
     */
    hotelId?: string;
    /**
     * Holds collection of articles.
     * @type {Array<ArticleInfoType>}
     * @memberof Articles
     */
    articles?: Array<ArticleInfoType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof Articles
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof Articles
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the Articles interface.
 */
export declare function instanceOfArticles(value: object): boolean;
export declare function ArticlesFromJSON(json: any): Articles;
export declare function ArticlesFromJSONTyped(json: any, ignoreDiscriminator: boolean): Articles;
export declare function ArticlesToJSON(value?: Articles | null): any;
