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
import type { CurrencyAmountType } from './CurrencyAmountType';
import type { RateCategoryType } from './RateCategoryType';
import type { RoomClassCodeInfoType } from './RoomClassCodeInfoType';
import type { SourceCodeInfoType } from './SourceCodeInfoType';
/**
 * Authorization configuration Rule Type.Authorization Rules specify the rule (or rules) that the property uses for credit card authorizations
 * @export
 * @interface AuthorizationConfigRuleType
 */
export interface AuthorizationConfigRuleType {
    /**
     * Property associated with authorization rule
     * @type {string}
     * @memberof AuthorizationConfigRuleType
     */
    hotelId?: string;
    /**
     * Authorization rule Room type.
     * @type {string}
     * @memberof AuthorizationConfigRuleType
     */
    roomType?: string;
    /**
     *
     * @type {RoomClassCodeInfoType}
     * @memberof AuthorizationConfigRuleType
     */
    roomClass?: RoomClassCodeInfoType;
    /**
     *
     * @type {SourceCodeInfoType}
     * @memberof AuthorizationConfigRuleType
     */
    sourceCode?: SourceCodeInfoType;
    /**
     *
     * @type {RateCategoryType}
     * @memberof AuthorizationConfigRuleType
     */
    rateCategory?: RateCategoryType;
    /**
     * Rate code associated to this rule.
     * @type {string}
     * @memberof AuthorizationConfigRuleType
     */
    rateCode?: string;
    /**
     * Guarantee code associated to this rule.
     * @type {string}
     * @memberof AuthorizationConfigRuleType
     */
    guaranteeCode?: string;
    /**
     * Rule Number for current authorization.
     * @type {number}
     * @memberof AuthorizationConfigRuleType
     */
    ruleNo?: number;
    /**
     * Rule Number description for current authorization.
     * @type {string}
     * @memberof AuthorizationConfigRuleType
     */
    ruleDescription?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof AuthorizationConfigRuleType
     */
    amount?: CurrencyAmountType;
    /**
     * Percentage to be applied to all rules with a percentage configured in the formula.
     * @type {number}
     * @memberof AuthorizationConfigRuleType
     */
    percentage?: number;
    /**
     * A maximum of two digits to set the number of days to authorize for a credit card. If the field is left blank, authorizations will be done for the entire stay duration of the reservation
     * @type {number}
     * @memberof AuthorizationConfigRuleType
     */
    maxDaysToAuthorize?: number;
    /**
     * Ignore deposits paid for calculation of the amount to be authorized.
     * @type {boolean}
     * @memberof AuthorizationConfigRuleType
     */
    ignoreDepositsYN?: boolean;
}
/**
 * Check if a given object implements the AuthorizationConfigRuleType interface.
 */
export declare function instanceOfAuthorizationConfigRuleType(value: object): boolean;
export declare function AuthorizationConfigRuleTypeFromJSON(json: any): AuthorizationConfigRuleType;
export declare function AuthorizationConfigRuleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizationConfigRuleType;
export declare function AuthorizationConfigRuleTypeToJSON(value?: AuthorizationConfigRuleType | null): any;
