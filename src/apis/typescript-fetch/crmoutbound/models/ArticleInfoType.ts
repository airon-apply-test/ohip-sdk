/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ArticlePostItType } from './ArticlePostItType';
import {
    ArticlePostItTypeFromJSON,
    ArticlePostItTypeFromJSONTyped,
    ArticlePostItTypeToJSON,
} from './ArticlePostItType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

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
export function instanceOfArticleInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ArticleInfoTypeFromJSON(json: any): ArticleInfoType {
    return ArticleInfoTypeFromJSONTyped(json, false);
}

export function ArticleInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ArticleInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': !exists(json, 'description') ? undefined : json['description'],
        'transactionCode': !exists(json, 'transactionCode') ? undefined : json['transactionCode'],
        'price': !exists(json, 'price') ? undefined : CurrencyAmountTypeFromJSON(json['price']),
        'postIt': !exists(json, 'postIt') ? undefined : ArticlePostItTypeFromJSON(json['postIt']),
        'universalProductCode': !exists(json, 'universalProductCode') ? undefined : json['universalProductCode'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'articleCode': !exists(json, 'articleCode') ? undefined : json['articleCode'],
        'inactive': !exists(json, 'inactive') ? undefined : json['inactive'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}

export function ArticleInfoTypeToJSON(value?: ArticleInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
        'transactionCode': value.transactionCode,
        'price': CurrencyAmountTypeToJSON(value.price),
        'postIt': ArticlePostItTypeToJSON(value.postIt),
        'universalProductCode': value.universalProductCode,
        'hotelId': value.hotelId,
        'articleCode': value.articleCode,
        'inactive': value.inactive,
        'orderSequence': value.orderSequence,
    };
}

